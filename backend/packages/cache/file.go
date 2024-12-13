package cache

import (
	"cdn-module/config"
	"cdn-module/packages/logger"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func ALLCache(config *config.CDN_CONFIG_JSON) *Safe {
	New := NewSafeMap()
	max := config.Cache_Size.Size()
	count := 0
	files, err := os.ReadDir(config.Cache_Dir)
	if err != nil {
		log.Fatalln(err)
	}
	for i, file := range files {
		start := time.Now()
		filePath := filepath.Join(config.Cache_Dir, file.Name())
		info, err := os.Stat(filePath)
		if err != nil {
			logger.ERR_Cache_FILE(file.Name(), i, len(files), New.MemoryUsage(), max, time.Since(start))
			fmt.Println(logger.Fg_BrightMagenta + "L  " + err.Error() + logger.Reset)
			continue
		}
		if max <= New.MemoryUsage() || max <= New.MemoryUsage()+info.Size() {
			logger.PASS_Cache_FILE(file.Name(), i, len(files), New.MemoryUsage()+info.Size(), max, time.Since(start))
			continue
		}
		data, err := os.ReadFile(filePath)
		if err != nil {
			logger.ERR_Cache_FILE(file.Name(), i, len(files), New.MemoryUsage()+info.Size(), max, time.Since(start))
			fmt.Println(logger.Fg_BrightMagenta + "L  " + err.Error() + logger.Reset)
			continue
		}
		New.Set(file.Name(), data)
		logger.LOAD_Cache_FILE(file.Name(), i, len(files), New.MemoryUsage(), max, time.Since(start))
		count++
	}
	logger.END_Cache(count, len(files)-count, New.MemoryUsage())
	return New
}

func (sm *Safe) Cache(config *config.CDN_CONFIG_JSON, start time.Time, file string) {
	max := config.Cache_Size.Size()
	count := 0
	files, err := os.ReadDir(config.Cache_Dir)
	if err != nil {
		log.Fatalln(err)
	}
	filePath := filepath.Join(config.Cache_Dir, file)
	info, err := os.Stat(filePath)
	if err != nil {
		logger.ERR_Cache_FILE(file, len(files)-1, len(files), sm.MemoryUsage(), max, time.Since(start))
		fmt.Println(logger.Fg_BrightMagenta + "L  " + err.Error() + logger.Reset)
		return
	}
	if max <= sm.MemoryUsage() || max <= sm.MemoryUsage()+info.Size() {
		logger.PASS_Cache_FILE(file, len(files)-1, len(files), sm.MemoryUsage()+info.Size(), max, time.Since(start))
		return
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.ERR_Cache_FILE(file, len(files)-1, len(files), sm.MemoryUsage()+info.Size(), max, time.Since(start))
		fmt.Println(logger.Fg_BrightMagenta + "L  " + err.Error() + logger.Reset)
		return
	}
	sm.Set(file, data)
	logger.UPLOAD_FILE(file, len(files)-1, len(files), sm.MemoryUsage(), max, time.Since(start))
	count++
}

func (sm *Safe) LoadFile(config config.CDN_CONFIG_JSON, file string) ([]byte, bool) {
	start := time.Now()
	caches, ex := sm.Get(file)
	if ex {
		logger.READ_FILE(file, true, sm.MemoryUsageKey(file), time.Since(start))
		return caches, true
	} else {
		filePath := filepath.Join(config.Cache_Dir, file)
		if info, err := os.Stat(filePath); os.IsNotExist(err) {
			logger.ERR_READ_FILE(file, false, 0, time.Since(start))
			fmt.Println(logger.Fg_BrightMagenta + "L  " + err.Error() + logger.Reset)
			return nil, false
		} else if err != nil {
			logger.ERR_READ_FILE(file, false, info.Size(), time.Since(start))
			fmt.Println(logger.Fg_BrightMagenta + "L  " + err.Error() + logger.Reset)
			return nil, false
		} else {
			data, err := os.ReadFile(filePath)
			if err != nil {
				logger.ERR_READ_FILE(file, false, info.Size(), time.Since(start))
				fmt.Println(logger.Fg_BrightMagenta + "L  " + err.Error() + logger.Reset)
				return nil, false
			}
			logger.READ_FILE(file, false, info.Size(), time.Since(start))
			return data, true
		}
	}
}

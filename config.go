package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Config struct {
	Level   string
	Threads uint8
	Port    int
}

func (c *Config) getIntFromArg(args []string, flag string, index int) (int, error) {
	var err error
	if len(args) == index {
		return 0, errors.New(fmt.Sprintf("missing %s argument value", flag))
	}
	argInt, err := strconv.Atoi(args[index+1])
	if err != nil {
		return 0, errors.New(fmt.Sprintf("didn't pass a correct argument to %s", flag))
	}
	return argInt, err
}

func (c *Config) Process(args []string) {
	for i, arg := range args {
		for _, flag := range FLAGS {
			if arg == flag {
				switch arg {
				case "-p":
					{
						argResult, err := c.getIntFromArg(args, "-p", i)
						if err != nil {
							panic(err)
						}
						c.Port = argResult
						continue
					}
				case "-t":
					{
						argResult, err := c.getIntFromArg(args, "-t", i)
						if err != nil {
							panic(err)
						}
						c.Threads = uint8(argResult)
						continue
					}
				case "--debug":
					{
						c.Level = DEBUG
						continue
					}
				}
			}
		}
	}
}

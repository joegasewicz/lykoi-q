package main

type Config struct {
	Level   string
	Threads uint8
	Port    int
}

func (c *Config) Process(args []string) {
	for i, arg := range args {
		for _, flag := range FLAGS {
			if arg == flag {
				switch arg {
				case "-p":
					{
						argResult, err := GetIntFromArg(args, "-p", i)
						if err != nil {
							panic(err)
						}
						c.Port = argResult
						continue
					}
				case "-t":
					{
						argResult, err := GetIntFromArg(args, "-t", i)
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

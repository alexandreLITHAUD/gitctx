package files

type Config struct {
	Aliases  map[string]string `ini:"aliases"`
	Merges   map[string]string `ini:"merges"`
	Branches map[string]string `ini:"branches"`
	Remotes  map[string]string `ini:"remotes"`
	Diff     map[string]string `ini:"diff"`
}

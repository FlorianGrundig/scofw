// TODO rename that struct
			// since we currently are not able to detect a commit we have to sync with current work tree manually
			gr.observer.UpdateCurrentScoSession()
func joinFilePaths(p1, p2 string) string {

	if p1 == "." {
		return p2
	}
	if p2 == "." {
		return p1
	}
	return filepath.Join(p1, p2)
}

// TODO this function should go into its own package because its the only function which is really tied to git
		log.Printf("No matching git change for file: %s", event.Name)
	var path string
	// under linux files in the root directory are reported as ./filename which is an invalid git tree path -> we have to remove the "./"
	if filepath.Dir(path) == "." {
		path = filepath.Base(event.Name)
	} else {
		path = event.Name
	}

	treeEntry, err := commitTree.EntryByPath(path)
	baseFolder := filepath.Join("diffs", joinFilePaths(gr.gitConfig.CurrentScoSession, filepath.Dir(event.Name)))
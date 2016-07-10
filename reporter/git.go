	"io/ioutil"
				lastChange, alreadyTracked := gr.gitConfig.LastChanges[event.Name]
					go gr.handleChange(event, lastChange)
			baseFolder := filepath.Join("diffs", gr.gitConfig.CurrentScoSession, filepath.Dir(event.Name))
			baseFile := filepath.Join(baseFolder, filepath.Base(event.Name))
			gr.util.CreateScoFolder(baseFolder)

			var contentA []byte
			var contentB []byte
			emptyContent := []byte("")
				contentA = blob.Contents()

				contentA = emptyContent
				contentB, err = ioutil.ReadFile(event.Name)
			} else {
				contentB = emptyContent // TODO is this really identical to delete?
			// TODO if event.Name referes to a new file -> the patch contains "new file mode 100644" -> we should change the file mode to the original settings
			patch, err := gr.repo.PatchFromBuffers(event.Name, event.Name, contentA, contentB, &options)
			defer patch.Free()
			_, err = patch.String()
			// TOOD use channel to publish change...
			// we store contentB as a snapshot of that file -> all further diffs will be made between workspace file and snapshot
			gr.util.WriteFile(&contentB, baseFile)
			gr.gitConfig.LastChanges[event.Name] = uint32(event.Op)
			gr.gitConfig.Persist()
			return
		}
	log.Printf("ERROR: No matching git change to file: %s", event.Op, event.Name)
func (gr *GitReporter) handleChange(event *fw.FileEvent, lastChange uint32) {
	log.Println("This is a change detected for", event.Name)

	options, err := git.DefaultDiffOptions()
	verifyNoError(err)

	// Specifying full patch indices.
	options.IdAbbrev = 40
	options.Flags |= git.DiffIncludeUntracked

	baseFolder := filepath.Join("diffs", gr.gitConfig.CurrentScoSession, filepath.Dir(event.Name))
	baseFile := filepath.Join(baseFolder, filepath.Base(event.Name))

	gr.util.CreateScoFolder(baseFile)

	var contentA *[]byte
	var contentB []byte
	emptyContent := []byte("")

	if event.Op == fw.Create || event.Op == fw.Write {
		contentA, err = gr.util.ReadScoFile(baseFile)
		verifyNoError(err)
	} else if event.Op == fw.Remove {
		// we create an empty file in diffs/.../a since this file event belongs to a new file
		contentA = &emptyContent
	} else {
		contentA = &emptyContent // TODO: how to handle renamed files? Maybe we should treat them as removed?
	}

	if event.Op != fw.Remove {
		contentB, err = ioutil.ReadFile(event.Name)
		verifyNoError(err)
	} else {
		contentB = emptyContent // TODO is this really identical to delete?
	}

	// TODO if event.Name referes to a new file -> the patch contains "new file mode 100644" -> we should change the file mode to the original settings
	patch, err := gr.repo.PatchFromBuffers(event.Name, event.Name, *contentA, contentB, &options)
	defer patch.Free()
	verifyNoError(err)
	patchString, err := patch.String()
	_, err = patch.String()
	verifyNoError(err)

	// TOOD use channel to publish change...
	log.Printf("\n%s", patchString)

	// we store contentB as a snapshot of that file -> all further diffs will be made between workspace file and snapshot
	gr.util.WriteFile(&contentB, baseFile)

	return

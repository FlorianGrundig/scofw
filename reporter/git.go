
	wkTree "github.com/floriangrundig/scofw/git"
	"github.com/floriangrundig/scofw/publisher"
	config                    *config.Config
	gitConfig                 *gitconfig.Config
	repo                      *git.Repository
	util                      *util.Util
	observer                  *wkTree.WorkTreeObserver
	fileEventChannel          chan *fw.FileEvent
	fileChangedMessageChannel chan *publisher.Message
func New(config *config.Config, gitConfig *gitconfig.Config, util *util.Util, observer *wkTree.WorkTreeObserver, fileEventChannel chan *fw.FileEvent, fileChangedMessageChannel chan *publisher.Message) *GitReporter {
		config:                    config,
		gitConfig:                 gitConfig,
		repo:                      repo,
		util:                      util,
		observer:                  observer,
		fileEventChannel:          fileEventChannel,
		fileChangedMessageChannel: fileChangedMessageChannel,
				break
			gr.observer.UpdateCurrentScoSession() // since we currently are not able to detect a commit we have to call update manually

				log.Printf("Received event %s %d (%s)\n", event.Name, event.Op, FileEventToString(event))
				lastChanges, anyChangesForSession := gr.gitConfig.LastChanges[gr.gitConfig.CurrentScoSession]

				var alreadyTracked bool
				var lastChange uint32

				if !anyChangesForSession {
					alreadyTracked = false
				} else {

					lastChange, alreadyTracked = lastChanges[event.Name]
				}
					gr.handleFirstChange(event)
					gr.handleChange(event, lastChange)
func FileEventToString(event *fw.FileEvent) string {

	if event.Op == fw.Chmod {
		return "chmod"
	}
	if event.Op == fw.Create {
		return "create"
	}
	if event.Op == fw.Write {
		return "write"
	}
	if event.Op == fw.Remove {
		return "remove"
	}
	if event.Op == fw.Rename {
		return "rename"
	}

	return "!!!unknown - this is not expected to happen!!!"
}

	baseFolder := filepath.Join("diffs", gr.gitConfig.CurrentScoSession, filepath.Dir(event.Name))
	baseFile := filepath.Join(baseFolder, filepath.Base(event.Name))
	gr.util.CreateScoFolder(baseFolder)

	var contentA []byte
	var contentB []byte
	emptyContent := []byte("")

	options.Flags |= git.DiffShowUntrackedContent
	var contentDeltaDetermined bool

				log.Println("This file is not tracked by git", event.Name)
			contentDeltaDetermined = true
			break

		}
	}

	if !contentDeltaDetermined {

		log.Printf("No matching git change for file: %s", event.Op, event.Name)

		_, err := commitTree.EntryByPath(event.Name)
		if err == nil {
			blob := gr.getOriginalBlob(commitTree, event)
			log.Printf("Anyway %s is tracked by git and git doesn't detect a change - assuming nothing has changed", event.Name)
			contentA = blob.Contents()
		} else {
			log.Printf("Going to fallback - assuming %s is a new file", event.Name)
			contentA = emptyContent

		contentB, err = ioutil.ReadFile(event.Name) // TODO can we be sure that this file is there (deleted?)?
		verifyNoError(err)
	}

	// TODO if event.Name referes to a new file -> the patch contains "new file mode 100644" -> we should change the file mode to the original settings
	patch, err := gr.repo.PatchFromBuffers(event.Name, event.Name, contentA, contentB, &options)
	defer patch.Free()
	verifyNoError(err)
	patchString, err := patch.String()
	_, err = patch.String()
	verifyNoError(err)

	// send to publisher
	gr.fileChangedMessageChannel <- &publisher.Message{
		FileEvent: event,
		Patch:     &patchString,
	}

	// we store contentB as a snapshot of that file -> all further diffs will be made between workspace file and snapshot
	gr.util.WriteFile(&contentB, baseFile)

	gr.storeLastChange(event)

}

func (gr *GitReporter) storeLastChange(event *fw.FileEvent) {
	lastChanges, anyChangesForSession := gr.gitConfig.LastChanges[gr.gitConfig.CurrentScoSession]

	if !anyChangesForSession {
		gr.gitConfig.LastChanges[gr.gitConfig.CurrentScoSession] = make(map[string]uint32)
	lastChanges, _ = gr.gitConfig.LastChanges[gr.gitConfig.CurrentScoSession]
	lastChanges[event.Name] = uint32(event.Op)
	gr.gitConfig.Persist()

	// Specifying full patch indices. TODO what is needed here?
	if event.Op&fw.Create == fw.Create || event.Op&fw.Write == fw.Write || event.Op&fw.Rename == fw.Rename { // TODO: how to handle renamed files? Maybe we should treat them as removed? (Beware -> IntelliJ stores the changes in a tmp file and renames that tmp file to the current file)
		log.Printf("Comparing current %s with last snapshot %s\n", event.Name, baseFile)
		contentA = &emptyContent // TODO this is not correct for IntelliJ -> when you revert your changes it's removed first and then created again... so we think it's a complete new file
		contentA = &emptyContent
	// publish event
	gr.fileChangedMessageChannel <- &publisher.Message{
		FileEvent: event,
		Patch:     &patchString,
	}
	gr.storeLastChange(event)
	"fmt"
	"os"
				// TODO if we receive a delete on a folder we have to deal with it -> e.g. deleting all files we know of (even files we don't know -> they not tracked but they should be listed in the git tracking...)

	result := ""
	if event.Op&fw.Chmod == fw.Chmod {
		result = fmt.Sprint("| chmod |", result)
	if event.Op&fw.Create == fw.Create {
		result = fmt.Sprint("| create |", result)
	if event.Op&fw.Write == fw.Write {
		result = fmt.Sprint("| write |", result)
	if event.Op&fw.Remove == fw.Remove {
		result = fmt.Sprint("| remove |", result)
	if event.Op&fw.Rename == fw.Rename {
		result = fmt.Sprint("| rename |", result)
	if result == "" {
		return "!!!unknown - this is not expected to happen!!!"
	}
	return result
			log.Printf("Going to fallback - %s seems to be an untracked file", event.Name)
		if event.Op&fw.Remove != fw.Remove {
			fmt.Println("not a removal")
			if _, err := os.Stat(event.Name); os.IsNotExist(err) {
				contentB = emptyContent
			} else {
				contentB, err = ioutil.ReadFile(event.Name) // TODO can we be sure that this file is there (deleted?)?
				verifyNoError(err)
			}
		} else {
			contentB = emptyContent
		}
	if event.Op&fw.Remove != fw.Remove {
		gr.util.WriteFile(&contentB, baseFile)
	} else {
		gr.util.RemoveFile(baseFile)
	}
	gr.storeLastChange(event)
	if lastChange&uint32(fw.Remove) != uint32(fw.Remove) { // TODO: how to handle renamed files? (Beware -> IntelliJ stores the changes in a tmp file and renames that tmp file to the current file)
	} else {
	if event.Op&fw.Remove != fw.Remove {

		if _, err = os.Stat(event.Name); os.IsNotExist(err) {
			contentB = emptyContent
		} else {
			contentB, err = ioutil.ReadFile(event.Name) // TODO can we be sure that this file is there (deleted?)?
			verifyNoError(err)
		}

	if event.Op&fw.Remove != fw.Remove {
		gr.util.WriteFile(&contentB, baseFile)
	} else {
		gr.util.RemoveFile(baseFile)
	}

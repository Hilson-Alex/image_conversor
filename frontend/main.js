const popError = window.go.conversor.App.PopError

const popMessage = window.go.conversor.App.PopMessage

function getFile() {
    window.go.conversor.App.GetFile()
        .then(filenames => {
            const fileText = document.getElementById('filename')
            fileText.textContent = setSelectedImages(filenames)
            enableButton(filenames)
        }).catch(error => {
            popError("Não foi possível ler as imagens", error)
        })
}

function convertImage() {
    const source = JSON.parse(document.getElementById('sourcePath').value)
    const destType = document.getElementById('dest-type').value
    
    const convertFunction =  source.length > 1 ? 
        () => window.go.conversor.App.ConvertMultiple(source, destType) :
        () => window.go.conversor.App.ConvertTo(source[0], destType) 

    console.log(convertFunction)
        
    convertFunction()
        .then(() => {
            popMessage("Salvo!", "Imagem salva com sucesso!")
        })
        .catch(error => {
            popError("Não foi possível salvar a imagem", error)
        })
}

function setSelectedImages (filenames) {
    document.getElementById('sourcePath').value = JSON.stringify(filenames)
    if (filenames.length === 1) {
        const filename = filenames[0]
        const shortname = filename.split('\\').pop()
        return `Imagem escolhida: ${shortname}`
    }
    if (filenames.length > 0) {
        return `${filenames.length} Imagens escolhidas`
    }
    return 'Nenhuma imagem escolhida'
}

function enableButton (filenames) {
    if (filenames.length > 0 ) {
        document.getElementById('convert').disabled = false
        return
    }
    document.getElementById('convert').disabled = true
}
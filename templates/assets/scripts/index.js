const svgIcon = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-copy" viewBox="0 0 16 16">
  <path fill-rule="evenodd" d="M4 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm2-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zM2 5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1h1v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1v1z"/>
</svg>`;

async function FadeOutContent() {
    const targetDiv = document.getElementById("welcome-text-content");
    let opacity;
    for (opacity = 1; opacity > 0; opacity -= 0.1) {
        await sleep(15);
        targetDiv.style.opacity = opacity;
    }
}

async function FadeInContent(contentID) {
    const targetDiv = document.getElementById(contentID);
    let opacity;
    for (opacity = 0; opacity <= 1; opacity += 0.1) {
        await sleep(20);
        targetDiv.style.opacity = opacity;
    }
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function copyRetrievedUrl() {
    const textToCopy = document.getElementById("text-to-copy");
    navigator.clipboard.writeText(textToCopy.textContent);
}

function removeSpinner() {
    const targetDiv = document.getElementById("url-loading-spinner");
    targetDiv.remove();
}

function putSpinner() {
    const spinnerHtml = `<div class="spinner-border text-info" role="status" id="url-loading-spinner"></div>`;
    const targetDiv = document.getElementById("url-loading-card");
    targetDiv.innerHTML = spinnerHtml
}

const TinifyUrlApi = async(urlAddress) => {
    const userObj = {
        url : urlAddress,
    }
    return fetch("/tinify", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userObj), 
        })
        .then(response => response.text())
        .then(text =>{
            return JSON.parse(text)
        });
  }
  

async function HandleUrlShortning() {
    FadeInContent("tinified-url-space");
    putSpinner();

    // make the HTTP call to the required endpoint
    const textAreaElement = document.getElementById("url-long-input");
    let urlInfo = textAreaElement.value;
    console.log("textArea Data : ", urlInfo);

    await TinifyUrlApi(urlInfo).then(urlData => {
        console.log("response : ",urlData);
        removeSpinner();
        const dataElement = `
        <p id="text-to-copy" class="col-10">`
        +urlData.url
        +`</p>
        <button class="btn btn-primary col-2" id="copy-button" onclick="copyRetrievedUrl()">`
        +svgIcon+`</button>`;
        const targetDiv = document.getElementById("url-loading-card");
        targetDiv.innerHTML = dataElement;
    })
}

async function EnterShortnerPage() {
    await FadeOutContent()
    const targetDiv = document.getElementById("welcome-text-content");
    const newHtml = `
        <div class="form-container ">
            <form id="shortenForm">
                <p class="fs-5 mt-3">Enter url to go-tiny...</p>
                <div class="mb-3">
                    <textarea class="form-control form-control-lg" form="testformid" name="url-input" id="url-long-input" cols="35" placeholder="Paste your long URL here..." required></textarea>
                </div>
                <br />
                <button id="shorten-url-button" type="submit" class="btn btn-custom btn-lg w-100" onclick="HandleUrlShortning() ">Shorten URL</button>
            </form>
        </div>
    `;
    if (targetDiv) {
        targetDiv.innerHTML = newHtml;
        FadeInContent("welcome-text-content");
    }
}
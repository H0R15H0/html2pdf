const sleep = (waitSeconds, callback) => {
	return new Promise(resolve => {
		setTimeout(() => {
			resolve(callback())
		}, waitSeconds * 1000)
	})	
}

const UNSELECTED = 'unselected'
const SELECTED = 'selected'

const buildSelectBox = (el) => {
  const domRect = el.getBoundingClientRect()
  const selectBox = document.createElement('div')
  selectBox.className = UNSELECTED
  selectBox.style.position = 'absolute'
  selectBox.style.top = `${document.documentElement.scrollTop + domRect.top}px`
  selectBox.style.left = `${domRect.left}px`
  selectBox.style.width = `${domRect.width}px`
  selectBox.style.height = `${domRect.height}px`
  selectBox.addEventListener('click', () => {
    if (selectBox.className == UNSELECTED) {
      chrome.storage.local.get('selectedUrls', (obj) => {
        obj.selectedUrls.push(el.href)
        chrome.storage.local.set({selectedUrls: obj.selectedUrls})
      })
      selectBox.className = SELECTED
    } else {
      chrome.storage.local.get('selectedUrls', (obj) => {
        chrome.storage.local.set({selectedUrls: obj.selectedUrls.filter((u) => { return u != el.href })})
      })
      selectBox.className = UNSELECTED
    }
  })
  return selectBox
}

const loadPage = () => {
  return sleep(1, () => {
    Object.values(document.getElementsByTagName("a")).forEach((a) => {
      const selectBox = buildSelectBox(a)
      document.body.appendChild(selectBox)
    })
  })
}

const init = () => {
  chrome.storage.local.clear()
  chrome.storage.local.set({selectedUrls: []})
}

chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "convert":
      init();
      loadPage().then(() => {console.log("load succeeded")})
      sleep(10, async () => {
        chrome.storage.local.get('selectedUrls', async (obj) => {
          const hrefs = obj.selectedUrls
          const resp = await fetch(`http://localhost:1323/users/4fae2802-7ad3-4ad3-a605-78685081cda8/pdfs`, {
            method: 'POST',
          })
          const pdf = await resp.json()
          await Promise.all(hrefs.map((url, idx) => {
            return fetch(`http://localhost:1323/users/4fae2802-7ad3-4ad3-a605-78685081cda8/pdfs/${pdf.id}/partials`, {
              method: 'POST',
              headers: { 'Content-Type': 'application/json' },
              body: JSON.stringify({
                number: idx,
                sourceHtmlUrl: url,
              })
            })
          }))
          sleep(20, async () => {
            const resp = await fetch(`http://localhost:1323/users/4fae2802-7ad3-4ad3-a605-78685081cda8/pdfs/${pdf.id}/unify`, {
              method: 'POST',
            })
            const blob = await resp.blob()
            const objUrl = URL.createObjectURL(blob);
    
            const link = document.createElement('a');
            link.href = objUrl;
            link.download = `${document.title}.pdf`;
            link.click();
          })
        })
      })
      break;
  
    default:
      break;
  }
});

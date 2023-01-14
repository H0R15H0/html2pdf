const sleep = (waitSeconds, callback) => {
	return new Promise(resolve => {
		setTimeout(() => {
			resolve(callback())
		}, waitSeconds * 1000)
	})	
}

const UNSELECTED = 'unselected'
const SELECTED = 'selected'
const MANAGER_MODAL_ID = 'hp_manager_modal'

const buildSelectBox = (aTag) => {
  aTag.className = UNSELECTED;
  aTag.addEventListener('click', (e) => {
    if (aTag.className == UNSELECTED) {
      // TODO: sort urls by y
      chrome.storage.local.get('selectedUrls', (obj) => {
        obj.selectedUrls.push(aTag.href)
        chrome.storage.local.set({selectedUrls: obj.selectedUrls})
      })
      aTag.className = SELECTED
    } else {
      chrome.storage.local.get('selectedUrls', (obj) => {
        chrome.storage.local.set({selectedUrls: obj.selectedUrls.filter((u) => { return u != aTag.href })})
      })
      aTag.className = UNSELECTED
    }
    e.preventDefault();
  })
}

const loadPage = () => {
  return sleep(1, () => {
    Object.values(document.getElementsByTagName("a")).forEach((a) => {
      buildSelectBox(a)
    })
  })
}

const buildManagerModal = () => {
  const managerModal = document.createElement('div')
  managerModal.id = MANAGER_MODAL_ID
  managerModal.style.position = 'fixed'
  managerModal.style.bottom = `0px`
  managerModal.style.right = `0px`
  const convertBtn = document.createElement('button')
  convertBtn.textContent = chrome.i18n.getMessage("endBtn")
  convertBtn.addEventListener('click', () => {
    chrome.storage.local.get('selectedUrls', (obj) => {
      fetchUrls(obj.selectedUrls).then(() => {
        // TODO: clear all elements
        console.log('fetch succeeded')
      })
    })
  })
  managerModal.appendChild(convertBtn)
  return managerModal
}

const init = () => {
  chrome.storage.local.clear()
  chrome.storage.local.set({selectedUrls: []})
  document.body.appendChild(buildManagerModal())
}

chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "convert":
      loadPage().then(() => {init()})
      break;
  
    default:
      break;
  }
});

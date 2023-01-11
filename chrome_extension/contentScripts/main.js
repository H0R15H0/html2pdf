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
      // TODO: sort urls by y
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

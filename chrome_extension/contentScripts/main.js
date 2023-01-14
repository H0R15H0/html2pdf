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

const aTagOnClick = (e) => {
  const aTag = e.path.find((path) => path.tagName == 'A')
  e.preventDefault();
  if (aTag.classList.contains(UNSELECTED)) {
    // TODO: sort urls by y
    chrome.storage.local.get('selectedUrls', (obj) => {
      obj.selectedUrls.push(aTag.href)
      chrome.storage.local.set({selectedUrls: obj.selectedUrls}, () => {
        aTag.classList.replace(UNSELECTED, SELECTED)
      })
    })
  } else {
    chrome.storage.local.get('selectedUrls', (obj) => {
      chrome.storage.local.set({selectedUrls: obj.selectedUrls.filter((u) => { return u != aTag.href })}, () => {
        aTag.classList.replace(SELECTED, UNSELECTED)
      })
    })
  }
}

const buildSelectBox = (aTag) => {
  aTag.classList.add(UNSELECTED);
  aTag.addEventListener('click', aTagOnClick)
}

const loadPage = () => {
  return sleep(1, () => {
    Object.values(document.getElementsByTagName("a")).forEach((a) => {
      buildSelectBox(a)
    })
  })
}

const clearSelectBox = (aTag) => {
  aTag.classList.remove(SELECTED, UNSELECTED)
  aTag.removeEventListener('click', aTagOnClick)
}

const clearPage = () => {
  Object.values(document.getElementsByTagName("a")).forEach((a) => {
    clearSelectBox(a)
  })
  document.getElementById(MANAGER_MODAL_ID).remove()
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
        console.log('fetch succeeded')
        clearPage()
      })
    })
  })
  managerModal.appendChild(convertBtn)
  const clearBtn = document.createElement('button')
  clearBtn.textContent = chrome.i18n.getMessage("clearBtn")
  clearBtn.addEventListener('click', () => {
    chrome.storage.local.clear()
    clearPage()
  })
  managerModal.appendChild(clearBtn)
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

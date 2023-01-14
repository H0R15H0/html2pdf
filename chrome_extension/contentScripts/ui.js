const UNSELECTED = 'unselected'
const SELECTED = 'selected'
const MANAGER_MODAL_ID = 'hp_manager_modal'

// ====================
// Select Box
// ====================
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

const clearSelectBox = (aTag) => {
  aTag.classList.remove(SELECTED, UNSELECTED)
  aTag.removeEventListener('click', aTagOnClick)
}

// ====================
// Manager Modal
// ====================
const buildConvertBtn = () => {
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
  return convertBtn
}

const clearPage = () => {
  Object.values(document.getElementsByTagName("a")).forEach((a) => {
    clearSelectBox(a)
  })
  document.getElementById(MANAGER_MODAL_ID).remove()
}

const buildClearBtn = () => {
  const clearBtn = document.createElement('button')
  clearBtn.textContent = chrome.i18n.getMessage("clearBtn")
  clearBtn.addEventListener('click', () => {
    chrome.storage.local.clear()
    clearPage()
  })
  return clearBtn
}

const buildManagerModal = () => {
  const managerModal = document.createElement('div')
  managerModal.id = MANAGER_MODAL_ID
  managerModal.style.position = 'fixed'
  managerModal.style.bottom = `0px`
  managerModal.style.right = `0px`
  managerModal.appendChild(buildConvertBtn())
  managerModal.appendChild(buildClearBtn())
  return managerModal
}

const initUI = () => {
  document.body.appendChild(buildManagerModal())
  Object.values(document.getElementsByTagName("a")).forEach((a) => {
    buildSelectBox(a)
  })
}

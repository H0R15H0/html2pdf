const convert = () => {
  chrome.tabs.query({active: true, lastFocusedWindow: true}, (tabs)=> {
    const currentTab = tabs[0].id
    chrome.tabs.sendMessage(currentTab, {name: "convert"}, () => {
    });
  })
}

const convertBtn = document.getElementById("convert")
convertBtn.textContent = chrome.i18n.getMessage("convert")
convertBtn.addEventListener('click', convert)


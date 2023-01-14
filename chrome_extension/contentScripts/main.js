const init = () => {
  chrome.storage.local.clear()
  chrome.storage.local.set({selectedUrls: []})
  initUI()
}

chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "convert":
      sleep(1, () => {
        init()
      })
      break;
  
    default:
      break;
  }
});

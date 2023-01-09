chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "convert":
      alert("convert ignite")
      break;
  
    default:
      break;
  }
});

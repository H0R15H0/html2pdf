const sleep = (waitSeconds, callback) => {
	return new Promise(resolve => {
		setTimeout(() => {
			resolve(callback())
		}, waitSeconds * 1000)
	})	
}

const loadPage = () => {
  return sleep(3, () => {
    const hrefs = Object.values(document.getElementsByTagName("a")).map((a) => a.href)
    return hrefs.slice(0, 3)
  })
}

chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "convert":
      loadPage().then(async (hrefs) => {
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
        return pdf
      }).then((pdf) => {
        sleep(20, async () => {
          const resp = await fetch(`http://localhost:1323/users/4fae2802-7ad3-4ad3-a605-78685081cda8/pdfs/${pdf.id}/unify`, {
            method: 'POST',
          })
          const blob = await resp.blob()
          const objUrl = URL.createObjectURL(blob);
  
          const link = document.createElement('a');
          link.href = objUrl;
          link.download = "unifiedPdf";
          link.click();
        })
      })
      break;
  
    default:
      break;
  }
});

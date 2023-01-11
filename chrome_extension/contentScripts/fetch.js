const fetchUrls = async (urls) => {
  const resp = await fetch(`http://localhost:1323/users/4fae2802-7ad3-4ad3-a605-78685081cda8/pdfs`, {
    method: 'POST',
  })
  const pdf = await resp.json()
  await Promise.all(urls.map((url, idx) => {
    return fetch(`http://localhost:1323/users/4fae2802-7ad3-4ad3-a605-78685081cda8/pdfs/${pdf.id}/partials`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        number: idx,
        sourceHtmlUrl: url,
      })
    })
  }))
  await sleep(20, async () => {
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
}

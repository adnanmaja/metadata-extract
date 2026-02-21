const API_BASE_URL = "http://127.0.0.1:8080/api"

export async function uploadFile(file) {
  try {
    const urlRes = await fetch(`${API_BASE_URL}/upload-url`, {
      method: 'POST'
    })

    if (!urlRes.ok) {
      throw new Error('Failed to get upload URL')
    }

    const { uploadUrl, blobUrl } = await urlRes.json()

    const uploadRes = await fetch(uploadUrl, {
      method: 'PUT',
      headers: {
        'x-ms-blob-type': 'BlockBlob',
        'Content-Type': file.type
      },
      body: file
    })

    if (!uploadRes.ok) {
      throw new Error('Blob upload failed')
    }

    const processRes = await fetch(`${API_BASE_URL}/upload`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        blobUrl,
        originalName: file.name
      })
    })

    if (!processRes.ok) {
      throw new Error('Processing failed')
    }

    return await processRes.json()
  } catch (error) {
    console.error('API error:', error)
    throw error
  }
}

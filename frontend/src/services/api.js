const API_BASE_URL = 'http://127.0.0.1:8080' 

export async function uploadFile(file) {
  const formData = new FormData()
  formData.append('file', file)
  
  try {
    console.log("Posting...")
    const response = await fetch(`${API_BASE_URL}/upload`, {
      method: 'POST',
      body: formData
    })
    
    if (!response.ok) {
      throw new Error('Upload failed')
    }
    
    return await response.json()
  } catch (error) {
    console.error('API error:', error)
    throw error
  }
}

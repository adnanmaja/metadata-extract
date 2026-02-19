import { ref } from 'vue'
import { uploadFile } from '../services/api'

export function useFileProcessor() {
  const isLoading = ref(false)
  const error = ref(null)
  const metadata = ref(null)

  async function processFile(file) {
    isLoading.value = true
    error.value = null
    try {
      console.log("Processing file:", file.name)
      const result = await uploadFile(file)
      metadata.value = result
      return result
    } catch (err) {
      error.value = err.message
      console.error('Failed to process file:', err)
    } finally {
      isLoading.value = false
    }
  }

  return {
    processFile,
    isLoading,
    error,
    metadata
  }
}
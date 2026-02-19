<script setup>
import { computed } from 'vue'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  metadata: {
    type: Object,
    required: true
  },
  category: {
    type: String,
    required: true,
    validator: (value) => ['basic', 'camera', 'location'].includes(value)
  }
})

const categoryMappings = {
  basic: [
    { key: 'FileName', label: 'File Name' },
    { key: 'FileSize', label: 'File Size (MiB)' },
    { key: 'FileType', label: 'File Type' },
    { key: 'DateTimeOriginal', label: 'Date Taken' },
    { key: 'ImageWidth', label: 'Width (px)' },
    { key: 'ImageHeight', label: 'Height (px)' },
    { key: 'Software', label: 'Software Used'}
  ],
  camera: [
    { key: 'Make', label: 'Camera Make' },
    { key: 'Model', label: 'Camera Model' },
    { key: 'LensInfo', label: 'Lens' },
    { key: 'FNumber', label: 'Aperture' },
    { key: 'ExposureTime', label: 'Shutter Speed' },
    { key: 'ISO', label: 'ISO' },
    { key: 'FocalLength', label: 'Focal Length' },
    { key: 'CameraSerial', label: 'Camera Serial Number'},
    { key: 'LensSerial', label: 'Lens Serial Number'}
  ],
  location: [
    { key: 'GPSLatitude', label: 'Latitude' },
    { key: 'GPSLongitude', label: 'Longitude' },
    { key: 'GPSAltitude', label: 'Altitude' },
    { key: 'City', label: 'City' },
    { key: 'Country', label: 'Country' }
  ]
}

const filteredMetadata = computed(() => {
  const mappings = categoryMappings[props.category]
  return mappings
    .map(({ key, label }) => ({
      label,
      value: props.metadata[key] || 'Not available'
    }))
    .filter(item => item.value !== undefined)
})
</script>

<template>
    <div class="result-section">
        <div class="section-label">{{ title }}</div>
        <div 
        v-for="(item, index) in filteredMetadata" 
        :key="index"
        class="result-rows"
      >
        <span class="metadata-label">{{ item.label }}:</span>
        <span class="metadata-value">{{ item.value }}</span>
        </div>
    </div>
</template>

<style scoped>

.result-section {
  background: #111;
  border: 1px solid #1e1e1e;
  border-radius: 12px;
  padding: 1.25rem 1.5rem;
}

.section-label {
  font-family: 'DM Mono', monospace;
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  color: #555;
  text-transform: uppercase;
  margin-bottom: 1rem;
}

.result-rows {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 0;
  border-bottom: 1px solid #1a1a1a;
  font-family: 'DM Mono', monospace;
  font-size: 0.85rem;
}

.result-rows:last-child {
  border-bottom: none;
}

.result-rows span:first-child { width: 30%; }
.result-rows span:last-child { width: 45%; }

</style>
<script setup>

</script>

<template>

    <div class="header">
      <div class="tag">EXIF · IPTC · XMP</div>
      <h1>Image<br><span class="accent">Metadata</span></h1>
      <p class="subtitle">Drop your image type shi.<br>See whats inside your image type shi.</p>
    </div>

</template>

<style scoped>
.header {
  text-align: center;
  max-width: 600px;
}

.tag {
  display: inline-block;
  font-family: 'DM Mono', monospace;
  font-size: 0.7rem;
  letter-spacing: 0.15em;
  color: #c8f060;
  border: 1px solid #c8f060;
  padding: 0.3rem 0.8rem;
  border-radius: 999px;
  margin-bottom: 1.2rem;
}

h1 {
  font-size: clamp(3rem, 8vw, 5.5rem);
  font-weight: 800;
  line-height: 1;
  letter-spacing: -0.03em;
  margin-bottom: 0.5rem;
}

.accent {
  color: #c8f060;
  display: block;
}

.subtitle {
  margin-top: 0.5rem;
  font-size: 1rem;
  color: #888;
  line-height: 1.6;
  font-weight: 400;
}
</style>
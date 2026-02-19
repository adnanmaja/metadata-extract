<script setup>
</script>

<template>
    <footer class="social-footer">
      <div class="social-container">
        <span class="footer-label">Rafa Atmaja</span>
        <div class="footer-social-links">

          <a href="https://github.com/adnanmaja" target="_blank" class="footer-social-link">
            <i class="fab fa-github"></i>
            <span>GitHub</span>
          </a>

          <a href="https://www.linkedin.com/in/rafa-adnan-0329b4321" target="_blank" class="footer-social-link">
            <i class="fab fa-linkedin"></i>
            <span>LinkedIn</span>
          </a>

        </div>
      </div>
    </footer>
</template>

<style scoped>

footer.social-footer {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 10;
  padding: 10px 0; 
  margin-top: auto;
}

.social-container {
  background: rgba(20, 20, 20, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  padding: 15px 25px;
  display: flex;
  align-items: center;
  gap: 20px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
  border: 1px solid #2a2d2e;
}

.footer-label {
  font-size: 14px;
  color: #9aa0a6;
  font-weight: 400;
}

.footer-social-links {
  display: flex;
  gap: 15px;
}

.footer-social-link {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 12px;
  background: rgba(40, 40, 40, 0.6);
  border-radius: 8px;
  text-decoration: none;
  color: #e8eaed;
  font-size: 13px;
  transition: all 0.2s;
  border: 1px solid transparent;
}

.footer-social-link:hover {
  background: #c8f060;
  border-color: #5f6368;
  transform: translateY(-2px);
}

.footer-social-link i {
  font-size: 16px;
}

.footer-social-link span {
  font-weight: 400;
}

.email-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  user-select: none;
}

.email-link {
  cursor: pointer;
}

.email-popup {
  position: absolute;
  bottom: calc(100% + 15px); 
  left: 50%;
  transform: translateX(-50%);
  background: rgba(20, 20, 20, 0.95);
  border: 1px solid #5f6368;
  border-radius: 8px;
  padding: 12px 16px;
  display: none;
  align-items: center;
  gap: 10px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.7);
  z-index: 100;
  white-space: nowrap; 
}


</style>
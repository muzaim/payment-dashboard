<script setup lang="ts">
const props = defineProps({
  show: {
    type: Boolean || undefined || String,
    default: false,
  },
})
</script>

<template>
  <div
    v-if="props.show"
    class="loading-overlay v-overlay__scrim"
  >
    <div class="loading-spinner">
      <div class="d-flex align-center gap-2">
        <span class="custom-loader">
          <VIcon icon="tabler-hourglass-low" />
        </span>
        <span class="loading-text">Loading.</span>
      </div>
    </div>
  </div>
</template>

  <style scoped>
  .loading-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5); /* Background transparan */
    z-index: 9999; /* Nilai z-index tinggi untuk berada di atas semua elemen */
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .loading-spinner {
    display: flex;
    flex-direction: column;
    align-items: center;
    color: white;
  }

  .icon-tabler-loader {
    animation: spin 1s linear infinite;
    stroke: white;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }

  p {
    margin-top: 10px;
    font-size: 18px;
  }
  .custom-loader {
    display: flex;
    animation: loader 1s infinite;
  }

  @keyframes loader {
    from {
      transform: rotate(0);
    }

    to {
      transform: rotate(360deg);
    }
  }

  .loading-text::after {
  content: '..';
  position: absolute;
  animation: dot-blink 1s steps(1) infinite;
}

@keyframes dot-blink {
  0% {
    visibility: hidden;
  }
  50% {
    visibility: visible;
  }
  100% {
    visibility: hidden;
  }
}
  </style>

<!--
  InfiniteSlideshow
  --
  This component contains an infinite slideshow. It has various props to modify its behavariour.
  Credits to https://github.com/biigpongsatorn/vue-infinite-slide-bar
  --
  Props:
    - delay: string - formatted with a trailing "s"
    - direction: string - either "normal" or "reverse"
    - duration: string - formatted with a trailing "s"
    - paused: boolean
-->

<template>
  <div class="is-container">
    <div class="is-element" :style="settings" data-cy="infinite-slideshow">
      <div class="is-section">
        <slot></slot>
      </div>
      <div class="is-section">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'InfiniteSlideshow',
  props: {
    delay: {
      type: String,
      default: '0s'
    },
    direction: {
      type: String,
      default: 'normal'
    },
    duration: {
      type: String,
      default: '12s'
    },
    paused: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    // Props are converted to CSS animation settings
    settings() {
      return {
        'animation-delay': this.delay,
        'animation-direction': this.direction,
        'animation-duration': this.duration,
        'animation-play-state': this.paused ? 'paused' : 'running'
      };
    }
  }
};
</script>

<style scoped>
@keyframes moveSlideshow {
  100% {
    transform: translateX(-50%);
  }
}

.is-container {
  width: 100%;
  overflow: hidden;
  padding: 15px;
}

.is-element {
  transform: translate3d(0, 0, 0); /* GPU intensive 😬 */
  position: relative;
  overflow: hidden;
  animation-name: moveSlideshow;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
  display: flex;
  width: max-content;
  min-width: 200%;
}

.is-section {
  width: 50%;
}
</style>

<template>
  <div class="content hidden-sm-and-down">
    <InfiniteSlideshow duration="32s" direction="reverse">
    <img
      height="250px"
      v-for="banner in banners"
      :key="banner.src"
      :src="banner.src"
      style="padding-left: 20px; padding-right: 20px;"
    >
    </InfiniteSlideshow>
  </div>
</template>

<script>
import InfiniteSlideshow from '@/components/InfiniteSlideshow';

export default {
  components: {
    InfiniteSlideshow,
  },
  data: () => ({
    banners: [
      { src: require('../../assets/banner-1.jpg') },
      { src: require('../../assets/banner-2.jpg') },
      { src: require('../../assets/banner-3.jpg') },
      { src: require('../../assets/banner-4.jpg') }],
  }),
};
</script>

<style lang="scss" scoped>
.content {
    flex-direction: column;
}
</style>

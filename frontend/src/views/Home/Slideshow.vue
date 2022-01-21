<template>
  <div class="content hidden-sm-and-down">
    <InfiniteSlideshow duration="32s" direction="reverse">
      <img
        height="250px"
        v-for="banner in banners"
        :key="banner.src"
        :src="banner.src"
        style="padding-left: 20px; padding-right: 20px;"
      />
    </InfiniteSlideshow>
  </div>
</template>

<script>
import InfiniteSlideshow from '@/components/InfiniteSlideshow';

export default {
  components: {
    InfiniteSlideshow
  },
  data: () => ({
    banners: [
      { src: require('../../assets/b1.png') },
      { src: require('../../assets/b2.png') },
      { src: require('../../assets/b3.png') },
      { src: require('../../assets/b4.png') },
      { src: require('../../assets/b5.png') },
      { src: require('../../assets/b6.png') },
      { src: require('../../assets/b7.png') },
      { src: require('../../assets/b8.png') },
      { src: require('../../assets/b9.jpg') }
    ]
  })
};
</script>

<style lang="scss" scoped>
.content {
  flex-direction: column;
}
</style>

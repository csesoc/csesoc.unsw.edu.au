<template>
  <div class="content">
    <InfiniteSlideshow duration="32s" direction="reverse">
    <img height="250px" src="@/assets/banner-1.jpg" style="padding-left: 20px; padding-right: 20px;">
    <img height="250px" src="@/assets/banner-2.jpg" style="padding-left: 20px; padding-right: 20px;">
    <img height="250px" src="@/assets/banner-3.jpg" style="padding-left: 20px; padding-right: 20px;">
    <img height="250px" src="@/assets/banner-4.jpg" style="padding-left: 20px; padding-right: 20px;">
    </InfiniteSlideshow>
    <InfiniteSlideshow duration="35s" direction="reverse">
    <img height="250px" src="@/assets/banner-1.jpg" style="padding-left: 20px; padding-right: 20px;">
    <img height="250px" src="@/assets/banner-2.jpg" style="padding-left: 20px; padding-right: 20px;">
    <img height="250px" src="@/assets/banner-3.jpg" style="padding-left: 20px; padding-right: 20px;">
    <img height="250px" src="@/assets/banner-4.jpg" style="padding-left: 20px; padding-right: 20px;">
    </InfiniteSlideshow>
  </div>
</template>

<script>
import InfiniteSlideshow from '@/components/InfiniteSlideshow';

export default {
  components: {
    InfiniteSlideshow,
  }
};
</script>

<style lang="scss" scoped>
.content {
    flex-direction: column;
}
</style>

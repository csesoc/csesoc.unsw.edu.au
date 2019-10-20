<template>
  <div>
    This is the post page.
    <BasePost />
    <ListComponent />
  </div>
</template>

<script>
import BasePost from '@/components/BasePost.vue';
import ListComponent from '@/components/ListComponent.vue';

export default {
  name: 'post',
  components: {
    BasePost,
    ListComponent,
  },
};
</script>

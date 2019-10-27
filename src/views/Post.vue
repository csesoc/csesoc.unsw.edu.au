<template>
  <div>
    <BasePost :parentData="myData" />
  </div>
</template>

<script>
import BasePost from '@/components/BasePost.vue';

export default {
  name: 'post',
  components: {
    BasePost,
  },
  data() {
    const title = 'Sample Header';
    const subtitle = 'Sample Subtitle';
    return {
      myData: {
        nav: title.toLowerCase(),
        header: title.toUpperCase(),
        subtitle: subtitle.toUpperCase(),
        posttime: '1st Jan 2019 11:59pm',
        // new lines need to be manually inserted so we need to parse them from the text
        text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
      },
    };
  },
};
</script>

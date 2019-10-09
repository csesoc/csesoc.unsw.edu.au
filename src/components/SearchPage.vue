<template>
  <div class="search">
    fjasdklfjsdakl;fjsadklf;j
    <input
      @input="search_text()"
      v-model="search.text"
      type="text"
      placeholder="Search..."
    />
  </div>
</template>
<script>
export default {
  name: 'SearchPage',
  el: '#searchpage',
  data() {
    const title = 'Header';
    return {
      nav: title.toLowerCase(),
      header: title.toUpperCase(),
      // new lines need to be manually inserted so we need to parse them from the text
      text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
    };
  },
};
</script>

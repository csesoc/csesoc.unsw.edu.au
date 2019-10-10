<template>
  <div>
    <p>This is a sample home view</p>
    <NavGrid :items="items"></NavGrid>
  </div>
</template>

<script>
import NavGrid from '@/components/NavGrid';

export default {
  data: () => ({
    apiUri: 'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    items: [],
  }),
  components: {
    NavGrid,
  },
  mounted() {
    fetch(this.apiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.items = responseJson;
      });
  },

};
</script>

<style scoped>
</style>

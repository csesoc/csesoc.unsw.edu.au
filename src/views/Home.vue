<template>
  <div>
    <p>This is a sample home view</p>
    <NavGrid :gridItems="gridItems"></NavGrid>
    <ListComponent/>
    <Footer/>
  </div>
</template>

<script>
import NavGrid from '@/components/NavGrid';
import ListComponent from '@/components/ListComponent';
import Footer from '@/components/Footer';



export default {
  data: () => ({
    gridApiUri: 'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    gridItems: [],
    listApiUri: 'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    listItems: [],
  }),

  components: {
    NavGrid,
    ListComponent,
    Footer,
  },

  mounted() {
    fetch(this.gridApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.gridItems = responseJson;
      });
  },

};
</script>

<style scoped>
</style>

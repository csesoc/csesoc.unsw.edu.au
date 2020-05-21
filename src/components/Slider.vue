<template>
  <div id="slider">
    <HeaderTitle id="slider__title" :title="'media'" />

    <section class="slider__items">
      <v-card
        v-for="item in items"
        :href="item.link ? item.link : '#0'"
        :key="item.id"
        class="ma-6 slider__item"
      >
        <v-lazy>
          <v-img :src="item.src" contain class="item__img"></v-img>
        </v-lazy>
      </v-card>
    </section>
  </div>
</template>

<style scoped>
#slider {
  font-family: "Quicksand", sans-serif;
  background: rgb(18, 76, 219);
  background: linear-gradient(
    125deg,
    rgba(18, 76, 219, 1) 0%,
    rgba(50, 112, 255, 1) 50%,
    rgba(30, 104, 255, 1) 100%
  );
  height: 80vh;
}

#slider__title {
  margin-left: 5vw;
  color: white;
  padding-top: 3vw;
}

.slider__items {
  display: flex;
  flex-direction: row;
  overflow-x: scroll;
  -webkit-overflow-scrolling: touch;
  -ms-overflow-style: -ms-autohiding-scrollbar;
  padding-bottom: 2vw;
  margin-left: 5vw;
  margin-right: 5vw;
  height: 60vh;
}

.slider__item {
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  flex: 0 0 auto;
  max-width: 30vw;
  background-color: rgba(0, 0, 0, 0.9);
}

.slider__item,
.item__img {
  height: 45vh;
}

@media screen and (max-width: 992px) {
  .slider__item,
  .item__img {
    max-width: 60vw;
    width: auto;
  }
}
.slider__item:hover {
  box-shadow: 0 5px 10px #0d24a6;
  transition: all 0.4s ease;
}

::-webkit-scrollbar {
  height: 12px;
}

::-webkit-scrollbar-track {
  border-radius: 10px;
  background: #465ff0;
}

::-webkit-scrollbar-thumb {
  border-radius: 10px;
  border: 1px solid #465ff0;
  background: #f7f8f9;
}
</style>

<script type="text/javascript">
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'Slider',
  components: { HeaderTitle },
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['title', 'items']
};
</script>

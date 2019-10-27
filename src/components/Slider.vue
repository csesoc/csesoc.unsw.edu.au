<template>
    <div id="slider">
        <h1 id="slider__title">#!/MEDIA</h1>
        <section class="slider__items">
            <v-card v-for="item in items" :key="item.title" :href="item.link" width="30vw" class="ma-6 slider__item">
              <v-img
                :src="item.src"
                class="white--text"
                height="180px"
                gradient="to bottom, rgba(0,0,0,.5), rgba(0,0,0,.6)"
              >
                <v-card-title class="justify-center fill-height align-center" v-text="item.title"></v-card-title>
              </v-img>
            </v-card>
        </section>
    </div>
</template>

<style scoped>
    @import url("https://fonts.googleapis.com/css?family=Quicksand");

    #slider{
        font-family: 'Quicksand', sans-serif;
        background-image: linear-gradient(#6278f5, #213bd1);
    }

    #slider__title{
        margin-left: 5vw;
        color: white;
        padding-top:  3vw;
    }

    .slider__items{
        display: flex;
        flex-direction: row;
        overflow-x: scroll;
        -webkit-overflow-scrolling: touch;
        -ms-overflow-style: -ms-autohiding-scrollbar;
        padding-bottom: 2vw;
        margin-left: 5vw;
        margin-right: 5vw;
    }

    .slider__item{
        box-shadow: 0 1px 2px rgba(0,0,0,.1);
        transition: all 0.3s ease;
        flex: 0 0 auto;
    }

    .slider__item:hover{
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
export default {
  name: 'Slider',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['items'],

};
</script>

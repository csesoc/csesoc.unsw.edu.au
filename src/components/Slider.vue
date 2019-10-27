<template>
    <div id="slider">
        <h1 id="slider__title">#!/MEDIA</h1>
    <section class="slider__items">
        <!-- <a v-for="item in items" class="slider__item"
        :href="item.link" :style="'background: url(' + item.src + '); background-size: cover; display: block;'">
            <div class="description">{{ item.title }}</div>
        </a> -->
        <v-card v-for="item in items" :href="item.link" width="30%" class="ma-4 slider__item">
          <v-img
            :src="item.src"
            class="white--text"
            height="180px"
            gradient="to bottom, rgba(0,0,0,.3), rgba(0,0,0,.8)"
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
    }

    #slider__title{
        margin-left: 5%;
        color: white;
        padding-top:  5%;
    }

    .slider__items{
        display: flex;
        flex-direction: row;
        overflow-x: scroll;
        -webkit-overflow-scrolling: touch;
        -ms-overflow-style: -ms-autohiding-scrollbar;
        padding-bottom: 2%;
        margin-left: 5%;
        margin-right: 5%;
        position: relative;
    }

    .slider__item{
        /* width: calc(80vw/3); */
        height: 10em;
        text-decoration:none;
        /* margin: 20px; */
        box-shadow: 0 1px 2px rgba(0,0,0,.1);
        transition: all 0.3s ease;
        flex: 0 0 auto;
    }

    #slider{
        background-image: linear-gradient(#6278f5, #213bd1);
    }

    .slider__item:hover{
        box-shadow: 0 5px 10px #0d24a6;
        transition: all 0.4s ease;
    }

    .description{
        padding: 10%;
        text-align: center;
        font-weight: bold;
        color:  #f7f8f9;
        font-size:  160%;
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

<template>
    <section class="nav__grid">
        <a v-for="item in items"  class="nav__box"
        :href="item.link"
        :style="'background: url(' + item.img + ');background-size: cover; display: block;'">
            <div class="nav__description">{{ item.description }}</div>
        </a>
    </section>
</template>


<style>
    @import url('https://fonts.googleapis.com/css?family=Quicksand');

    .nav__grid{
        font-family: 'Quicksand', sans-serif;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        padding-bottom: 2%;
        margin-left: 5%;
        margin-right: 5%;
        position: relative;
    }

    .nav__box{
        width: calc(80%/3);
        height: 10em;
        text-decoration:none;
        margin: 20px;
        box-shadow: 0 1px 2px rgba(0,0,0,.1);
        transition: all 0.3s ease;
        flex: 0 0 auto;
    }

    .nav__box:hover{
        box-shadow: 0 5px 10px #777;
        transition: all 0.4s ease;
    }

    .nav__description{
        padding: 10%;
        text-align: center;
        font-weight: bold;
        color:  #f7f8f9;
        font-size:  160%;
    }

</style>

<script type="text/javascript">
export default {
    name: 'NavGrid',
    data: () => (
        {
          items: []
          // items have description, image url, and link
        }
    ),

    mounted() {
        fetch("")
        .then (r => { return r.json() })
        .then(responseJson => {
            console.log(responseJson)
            this.items = responseJson
        })
    }

};
</script>

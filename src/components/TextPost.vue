<template>

  <div>
  <v-container class="pa-md-10 mx-lg-auto">
    <!-- maybe a loop to print out path to home-->
    <div class="links"> <span class="underline">home</span> /
     <span class="underline">{{nav}}</span> </div>
<v-flex xs12>
<div id="nav">


<section class="nav__grid">
    <a v-for="item in items" class="grid" v-bind:key="item.label"
    :href="item.link" :style="'background: url(' + item.img +
    '); background-size: cover; display:block;'">
        <div class="description">{{ item.label }}</div>
    </a>
</section>
</div>

  <h1> #!/{{header}} </h1>
      </v-flex>
</v-container>
          <img src="../assets/fbcoverpic.png" alt="sample"
            height = "350px"
              width=100%
              object-fit= contain>


<v-container class="pa-md-10 mx-lg-auto">
  <p> {{text}} </p>
  </v-container>
    </div>
</template>

<script>
export default {
  name: 'TextPost',
  el: '#nav',
  data() {
    const title = 'Header';
    return {
      nav: title.toLowerCase(),
      header: title.toUpperCase(),
      // new lines need to be manually inserted so we need to parse them from the text
      text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
      items: [
        { label: 'edwe', img: '../assets/fbcoverpic.png', link: '' },
        { label: 'wedfw', img: '../assets/fbcoverpic.png', link: '' },
        { label: 'wefwef', img: '../assets/fbcoverpic.png', link: '' },
      ],
    };
  },
};
</script>

<style scoped>

@import url('https://fonts.googleapis.com/css?family=Quicksand&display=swap');

h1 {
  font-size:48px;
  font-weight: normal;
}

body {
  font-family: 'Quicksand', sans-serif;
  text-align: left;
}
p {
    font-size:14px;
    font-weight: 300;
    -webkit-font-smoothing: subpixel-antialiased;

}

p span {
    font-size:16px;
    font-style: italic;
    display:inline-block;
}

links {
  font-size: 20px;
}

.underline {
  border-bottom: 1px solid grey;
}

/* grid */
.nav__grid{
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        padding-bottom: 2%;
        margin-left: 5%;
        margin-right: 5%;
        position: relative;
    }

    .grid{
        width: calc(80%/3);
        height: 10em;
        text-decoration:none;
        margin: 20px;
        box-shadow: 0 1px 2px rgba(0,0,0,.1);
        transition: all 0.3s ease;
        flex: 0 0 auto;
    }

    .grid:hover{
        box-shadow: 0 5px 10px #777;
        transition: all 0.4s ease;
    }

    .description{
        padding: 10%;
        text-align: center;
        font-weight: bold;
        color:  #f7f8f9;
        font-size:  160%;
    }

</style>


<template>
    <div>
       <!-- <svg class="defs-only" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
       <symbol id="wave">
       <svg viewBox="0 0 100 50" preserveAspectRatio="none" >
          <g>
        	 <path d="M100,30 Q70,40 50,30 T0,30 v20 h100Z"
              style="fill:#4a4a4a; stroke-linejoin:round; stroke:#ffffff" stroke-width="0"/>
          </g>
       </svg>
       </symbol>
       </svg>
          <svg id="top-wave" width="50" height="50" style="width: 100%; height: 350px">
          <use xlink:href="#wave"/>
       </svg> -->

    <v-footer
      dark
      padless
      id="tiers"
    >

        <v-card
          id="tiers__sponsors"
          width="100%"
          class="white--text text-center"
        >

        <v-card-text class="white--text">

          <section class="tiers__grid">
              <a v-for="sponsor in tierOne" class="tiers__box" :href="sponsor.link" :key="sponsor">
                  <img :src="sponsor.logo" style=" max-width:300px;max-height:105px;">
              </a>
          </section>
          <section class="tiers__grid">
              <a v-for="sponsor in tierTwo" class="tiers__box" :href="sponsor.link" :key="sponsor">
                  <img :src="sponsor.logo" style="max-width:200px;max-height:75px">
              </a>
          </section>
          <section class="tiers__grid">
              <a v-for="sponsor in tierThree" class="tiers__box" :href="sponsor.link" :key="sponsor">
                  <img :src="sponsor.logo" style="max-width:100px;max-height:50px">
              </a>
          </section>

        </v-card-text>

        <h2>CONTACT US</h2>
        <v-card-text class="white--text text-center">
            <a href="https://www.facebook.com/csesoc">Facebook Page</a>
            | <a href="https://www.facebook.com/groups/csesoc">Facebook Group</a>
            | <a href="https://www.linkedin.com/company/csesoc/about/">LinkedIn</a>
            <br>
                    B03 CSE Building K17, UNSW
            <br>
            <a id='tiers__link' href="mailto:csesoc@csesoc.org.au">csesoc@csesoc.org.au</a>
        </v-card-text>


        <v-divider></v-divider>

        <v-card-text class="white--text text-center">
          {{ new Date().getFullYear() }} — <strong>CSESoc</strong>
        </v-card-text>
      </v-card>
    </v-footer>
    </div>
</template>


<style scoped>

    @import url("https://fonts.googleapis.com/css?family=Quicksand");

    #tiers {
        margin-top: 0px;
        /* border-top-right-radius: 5% 10%;
        border-top-left-radius: 5% 10%; */
        font-family: "Quicksand", sans-serif;
    }
    #top-wave {
        margin: 0;

    }

    #tiers__sponsors {
        background: #4a4a4a;
        padding-top:  5%;
    }

    .white--text {
        font-size: 16px;
        line-height: 28px;
    }
    .white--text a {
        color: white;
    }

    .white--text a:hover {
        color: #427bff;
    }

    .tiers__grid{
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
        margin-left: 1%;
        margin-right: 1%;
        position: relative;
    }


    .tiers__box{
        margin: 20px;
    }


</style>

<script type="text/javascript">


export default {
  name: 'Footer',
  data: () => (
    {
      sponsors: [
        {
          name: 'google', logo: 'https://spinmodern.com/wp-content/uploads/2015/09/google-logo-white.png', tier: 1000, link: 'http://google.com',
        },
        {
          name: 'microsoft', logo: 'https://www.lanmark.com/wp-content/uploads/2018/10/Microsoft-Logo-White.png', tier: 100, link: '',
        },
        {
          name: 'atlassian', logo: 'https://cdn.shortpixel.ai/client/to_webp,q_glossy,ret_img/https://www.generativ.ch/wp-content/uploads/2019/02/Atlassian-vertical-white@2x-rgb.png', tier: 100, link: '',
        },
        {
          name: 'canva', logo: 'https://static.canva.com/static/images/canva_logo_100x100@2x.png', tier: 10, link: '',
        },
        {
          name: 'deloitte', logo: 'https://solsticemedia.com.au/wp-content/uploads/2018/03/Deloitte-logo-white.png', tier: 10, link: '',
        },
        {
          name: 'amazon', logo: 'https://golpc.com/wp-content/uploads/2018/11/amazon-logo-white.png', tier: 10, link: '',
        },
        {
          name: 'jane street', logo: 'https://opensource.janestreet.com/assets/JS_logo-d7838b558a1de6c51553426ab5a2bba474510c41c6a5e910a9e30524a32dec27.png', tier: 1000, link: '',
        },
        {
          name: 'google', logo: 'https://spinmodern.com/wp-content/uploads/2015/09/google-logo-white.png', tier: 1000, link: 'http://google.com',
        },
        {
          name: 'microsoft', logo: 'https://www.lanmark.com/wp-content/uploads/2018/10/Microsoft-Logo-White.png', tier: 100, link: '',
        },
        {
          name: 'atlassian', logo: 'https://cdn.shortpixel.ai/client/to_webp,q_glossy,ret_img/https://www.generativ.ch/wp-content/uploads/2019/02/Atlassian-vertical-white@2x-rgb.png', tier: 100, link: '',
        },
        {
          name: 'canva', logo: 'https://static.canva.com/static/images/canva_logo_100x100@2x.png', tier: 10, link: '',
        },
        {
          name: 'deloitte', logo: 'https://solsticemedia.com.au/wp-content/uploads/2018/03/Deloitte-logo-white.png', tier: 10, link: '',
        },
        {
          name: 'amazon', logo: 'https://golpc.com/wp-content/uploads/2018/11/amazon-logo-white.png', tier: 10, link: '',
        },
        {
          name: 'jane street', logo: 'https://opensource.janestreet.com/assets/JS_logo-d7838b558a1de6c51553426ab5a2bba474510c41c6a5e910a9e30524a32dec27.png', tier: 1000, link: '',
        },
        {
          name: 'commbank', logo: 'https://www.vml.com/australia/sites/vml.australia/files/client/logo/cba-logo.png', tier: 100, link: '',
        },
      ],
    }
  ),
  computed: {
    tierOne() {
      return this.sponsors.filter(x => x.tier >= 1000);
    },
    tierTwo() {
      return this.sponsors.filter(x => (x.tier >= 100 && x.tier < 1000));
    },
    tierThree() {
      return this.sponsors.filter(x => (x.tier >= 10 && x.tier < 100));
    },

  },
  mounted() {
    fetch('')
      .then(r => r.json())
      .then((responseJson) => {
        console.log(responseJson);
        this.sponsors = responseJson;
      });
  },
};


</script>

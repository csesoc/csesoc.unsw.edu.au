<template>
  <div class="content">
    <v-app>
        <!-- Sizes your content based upon application components -->
        <v-main>
            <!-- Provides the application the proper gutter -->
            <v-container fluid>
                <div>
                    <h1>Principal Sponsors</h1>
                    <div v-for="sponsor in tierOne" :key="sponsor.id">
                        <a :href="sponsor.url">
                            <img :src="sponsor.logo" style=" max-width:300px;max-height:105px;" />
                        </a>
                        <p>{{sponsor.detail}}</p>
                        <v-divider/>
                    </div>
                    <h1>Major Sponsors</h1>
                    <div v-for="sponsor in tierOne" :key="sponsor.id">
                        <a :href="sponsor.url">
                            <img :src="sponsor.logo" style=" max-width:300px;max-height:105px;" />
                        </a>
                        <p>{{sponsor.detail}}</p>
                        <v-divider/>
                    </div>
                    <h1>Affiliate Sponsors</h1>
                    <div v-for="sponsor in tierOne" :key="sponsor.id">
                        <a :href="sponsor.url">
                            <img :src="sponsor.logo" style=" max-width:300px;max-height:105px;" />
                        </a>
                        <p>{{sponsor.detail}}</p>
                        <v-divider/>
                    </div>
                </div>
            </v-container>
        </v-main>
    </v-app>
  </div>
</template>

<script type="text/javascript">
import APIClient  from '../utils/APIClient'
export default {
  name: 'Footer',
  data: () => ({
    // sponsors have name, logo, 'tier' value, and link
    sponsors: []
  }),
computed: {
    // functions to determine sizing category of sponsor based on their value
    tierOne() {
      return this.sponsors.filter(x => x.tier == 2);
    },
    tierTwo() {
      return this.sponsors.filter(x => x.tier == 1);
    },
    tierThree() {
      return this.sponsors.filter(x => x.tier == 0);
    }
  },
  mounted() {
      APIClient.sponsorsAPI()
      .then((responseJson) => {
        this.sponsors = responseJson;
      });
  }
};
</script>

<style scoped>
</style>
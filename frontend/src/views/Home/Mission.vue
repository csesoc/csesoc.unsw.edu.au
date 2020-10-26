<!--
  Mission
  --
  Mission statement of CSESOC
-->

<template>
   <div id="mission" class="content">
    <v-container>
    <HeaderTitle :title="'CSESOC\'S MISSION'" />
    <v-row no-gutters justify="end">
      <v-col xs="12" sm="6">
        <v-img contain :src="mission.image" class="image"/>
      </v-col>
      <v-col xs="12" sm="6" align-self="center">
        <p class="text--md">
          {{mission.statement}}
        </p>
      </v-col>
    </v-row>
    </v-container>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'Mission',
  components: {
    HeaderTitle
  },
  data: () => ({
    mission: {
      image: require('../../assets/landing_page/mission.png'),
      statement: '"To empower every person and every organization on the planet to achieve more. We strive to create local opportunity, growth, and impact in every country around the world."',
    },
  }),
};
</script>
<style lang="scss" scoped>
#mission {
  // background-image: linear-gradient($dark-color-1, transparent, $dark-color-1), url(../../assets/mission.jpg);
  // height: 100vh;
  // background-size: 50%;
  // padding-top: $space-sm;
}

.image {
  max-width: 100%;
  max-height: 100%;
}
</style>

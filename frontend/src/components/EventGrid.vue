<!--
  EventGrid
  --
  This component contains a grid of event tiles containing basic information regarding each event.
  It is currently being used in the landing page.
  --
  Props:
    - events: list of event objects - containing the fields: id, name, start_time, fb_event_id, fb_cover_img
-->

<template>
  <v-container fluid>
    <!-- Catch a lack of events. -->
    <div v-if="events.length == 0">Stay tuned to our Facebook page for upcoming events!</div>
    <v-row class="justify-center">
      <v-col sm="8" md="6" lg="6" xl="4" v-for="event in events" :key="event.id">
        <!-- <v-card style = "position: relative" class="ma-1 card" height="100%"  target="_blank">
          <v-img
            :src="event.fb_cover_img"
            class="grid-img"
            height = "150px"
            gradient="to bottom, rgba(255,255,255,0.1) 0%, rgba(41,41,41,0.6) 50%, rgba(24,24,24,0.8) 100%"
          >
          </v-img>
          <v-card-text height="100%">
            <div class="name" v-text="event.name"/>
            <div class="event__place" v-if="event.place != undefined" v-text="'@ '+event.place"/>
            <div class="description" v-text="event.description"/>
          </v-card-text>
          <div class="date">
            <div class="day" v-text ="new Date(event.start_time*1000).getDate()"></div>
            <div class="month" v-text="getMonthString(event.start_time)"></div>
          </div>
          <v-card-actions>
            <a :href="'https://facebook.com/' + event.fb_event_id"><div class="link">Learn more ▶</div></a>
          </v-card-actions>

        </v-card> -->
      <Event :event = 'event'></Event>

      </v-col>
    </v-row>
  </v-container>
</template>

<script type="text/javascript">
import Event from '@/components/Event';

export default {
  name: 'EventGrid',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['events'],
  components: {
    Event
  }
};
</script>

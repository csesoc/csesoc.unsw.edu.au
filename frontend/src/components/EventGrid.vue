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
  <v-container>
    <!-- Catch a lack of events. -->
    <div v-if="events.length == 0">Stay tuned to our Facebook page for upcoming events!</div>
      <v-slide-group show-arrows class = "hidden-sm-and-down">
        <v-slide-item v-for="event in events" :key="event.id">
          <Event :event = 'event'></Event>
        </v-slide-item>
      </v-slide-group>
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

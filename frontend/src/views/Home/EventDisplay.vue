<!--
  EventDisplay
  --
  This component contains a grid of event tiles containing basic information regarding each event.
  It is currently being used in the landing page.
  --
  Props:
    - events: list of event objects - containing the fields: id, name, start_time, fb_event_id, fb_cover_img
    - updated: the time since the JSON file of events was last updated.
-->

<template>
  <div id="event-display" class="content">
    <v-container data-cy="event-section">
      <HeaderTitle
        title="UPCOMING EVENTS"
        subtitle="We run a wide-variety of events for fun, learning new skills and careers."
      />
      <!-- Catch a lack of events, or if events haven't been updated in 60 days. -->
      <div data-cy="event-alert" v-if="events.length == 0 | updated > 86400 * 1000 * 60">
        For full listings, check out the CSESoc Discord or our Facebook page!
      </div>

      <!-- Display all events in a sliding component on desktop viewports. -->
      <v-slide-group data-cy="event-slider" show-arrows dark class="hidden-sm-and-down">
        <v-slide-item v-for="event in events" :key="event.id" style="margin-right: 20px;">
          <Event :event = "event"></Event>
        </v-slide-item>
      </v-slide-group>

      <!-- Otherwise, list all events on mobile. -->
      <v-row data-cy="event-list" class="hidden-md-and-up" v-for="event in events.slice(0,3)" :key="event.id">
          <Event :event = "event" style="margin: 0 auto; margin-bottom: 20px;"></Event>
      </v-row>
    </v-container>
  </div>
</template>

<script type="text/javascript">
import Event from '@/components/Event';
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'EventDisplay',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['events', 'updated'],
  components: {
    Event,
    HeaderTitle,
  }
};
</script>

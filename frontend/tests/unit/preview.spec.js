// Libraries
import Vuetify from 'vuetify';

// Components
import Preview from '@/components/Preview'

// Utilities
import { mount, createLocalVue } from '@vue/test-utils';

describe('Preview.vue', () => {
  let localVue;
  let vuetify;
  let wrapper;

  beforeEach(() => {
    localVue = createLocalVue();
    vuetify = new Vuetify();
    wrapper = mount(Preview, {
      localVue,
      vuetify
    });
  });

  it('should load resources as components', () => {
    const titleArray = wrapper.findAll('.v-list-item__title');
    const subtitleArray = wrapper.findAll('.v-list-item__subtitle');
    // see if it exists and resources are being fetched
    expect(titleArray.length).toEqual(4);
    expect(subtitleArray.length).toEqual(4);
    // see if every list item has an action
  });

  it('should allow for on hover that changes the preview', () => {

  });

  it('should preview the hovered over resource', () => {

  });
});

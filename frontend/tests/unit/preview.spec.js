// Libraries
import Vuetify from 'vuetify';

//Components
import Preview from '@/components/Preview'

//Utilities
import { mount, createLocalVue } from '@vue/test-utils';

describe('Preview.vue', () => {
  let localVue, vuetify, wrapper;

  beforeEach(() => {
    localVue = createLocalVue();
    vuetify = new Vuetify();
    wrapper = mount(Preview, {
      localVue,
      vuetify
    });
  });

  it('should load resources as components', () => {
    
  });

  it('should allow for on hover that changes the preview', () => {

  });

  it('should preview the hovered over resource', () => {

  })
})
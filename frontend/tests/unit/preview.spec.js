// Libraries
import Vuetify from 'vuetify';

// Components
import Preview from '@/components/Preview'

// Utilities
import { mount, createLocalVue } from '@vue/test-utils';

// Testing data
const data = require('../../../backend/static/resource.json')

describe('Preview.vue', () => {
  let localVue;
  let vuetify;
  let wrapper;

  beforeEach(() => {
    localVue = createLocalVue();
    vuetify = new Vuetify();
    wrapper = mount(Preview, {
      propsData: {
        items: data,
      },
      localVue,
      vuetify,
    });
  });

  it('should load resources as components with title, description and link', () => {
    const itemArray = wrapper.findAll('.v-list-item');
    //see if the resources have been fetched
    expect(itemArray.length).toEqual(4);
    itemArray.wrappers.forEach(item => {
      //all item should have a title and description related to it
      expect(item.attributes('href')).not.toMatch('^$');
      item.get('.text-h4');
      item.get('.text-subtitle-1');
    })
  });
  
  it('should load an image from the fetch in the preview', () => {
    const img = wrapper.get('.preview-img');
    img.get('.v-image__image');
  });
  
  it('should allow for on hover that changes the preview', () => {
    const itemArray = wrapper.findAll('.v-list-item');
    //still figuring this out.
  })
});

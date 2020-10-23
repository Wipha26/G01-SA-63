import { createPlugin } from '@backstage/core';
import Dispense from './components/Dispense';
import login from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/dispense', Dispense);
    router.registerRoute('/', login);

  },
});

import React from 'react';
import { Route, Switch, MemoryRouter } from 'react-router-dom';
import { rest } from 'msw'
import { setupServer } from 'msw/node'
import { render, fireEvent, waitFor, screen, getByTestId, within, queryByTestId } from '@testing-library/react'
import '@testing-library/jest-dom/extend-expect'
import LandingPage from './components/views/LandingPage/LandingPage';
import RightMenu from './components/views/NavBar/Sections/RightMenu';

  
test('Landing Page join goes to Register', async () => {
  const mockHistoryPush = { push: jest.fn() };
  //set json object to storage 
  render(
    <MemoryRouter initialEntries={[`/`]}>
      <LandingPage history={mockHistoryPush}/>
    </MemoryRouter>
  );
  fireEvent.click(screen.getByRole ('button'));
  expect(mockHistoryPush.push.mock.calls[0]).toEqual([ '/register'])
    //remove object
}); 

test('Right Menu goes to Register', async () => {
  const mockHistoryPush = { push: jest.fn() };
  //set json object to storage 
  render(
    <MemoryRouter initialEntries={[`/`]}>
      <RightMenu history={mockHistoryPush}/>
    </MemoryRouter>
  );
  screen.debug()
  fireEvent.click(screen.getByText ('Profile'));
  console.log(mockHistoryPush.push.mock)
  //expect(mockHistoryPush.push.mock.calls[0]).toEqual([ '/register'])
    //remove object
}); 
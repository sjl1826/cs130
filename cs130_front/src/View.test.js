import React from 'react';
import { Route, Switch, MemoryRouter } from 'react-router-dom';
import { rest } from 'msw'
import { setupServer } from 'msw/node'
import { render, fireEvent, waitFor, screen, getByTestId, within, queryByTestId } from '@testing-library/react'
import '@testing-library/jest-dom/extend-expect'
import LandingPage from './components/views/LandingPage/LandingPage';
import RegisterPage from './components/views/RegisterPage/RegisterPage';
import { USER_SERVER } from './Config';

const server = setupServer(
  rest.post(`/api/v1/user/register`, (req, res, ctx) => {
    return res(ctx.json({ success: true }))
  })
)

beforeAll(() => server.listen())
afterEach(() => server.resetHandlers())
afterAll(() => server.close())


test('Landing Page join goes to Register', async () => {
  const mockHistoryPush = { push: jest.fn() };
  render(
    <MemoryRouter initialEntries={[`/`]}>
      <LandingPage history={mockHistoryPush}/>
    </MemoryRouter>
  );
  fireEvent.click(screen.getByRole('button'));
  expect(mockHistoryPush.push.mock.calls[0]).toEqual([ '/register'])
}); 

test('Register Page goes to Login Page', async () => {
  const mockHistoryPush = { push: jest.fn() };
  render(
    <MemoryRouter initialEntries={[`/register`]}>
      <RegisterPage history={mockHistoryPush}/>
    </MemoryRouter>
  );
  fireEvent.click(screen.getByRole('button'))
  waitFor(() => screen.debug());
  //console.log(mockHistoryPush.push.mock);
  //expect(mockHistoryPush.push.mock.calls[0]).toEqual([ '/login'])
}); 
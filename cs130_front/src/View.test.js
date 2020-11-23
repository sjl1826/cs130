import React from 'react';
import { Route, Switch, MemoryRouter } from 'react-router-dom';
import { rest } from 'msw'
import { setupServer } from 'msw/node'
import { render, fireEvent, waitFor, screen } from '@testing-library/react'
import '@testing-library/jest-dom/extend-expect'
import LandingPage from './components/views/LandingPage/LandingPage';
import ShortListing from './components/views/ProfilePage/ShortListing';
import MyListings from './components/views/ProfilePage/MyListings';
import CourseView from './components/views/ProfilePage/CourseView';
import SchedulerPage from './components/views/SchedulerPage/SchedulerPage';

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

  await fireEvent.click(screen.getByRole('button'));
  expect(mockHistoryPush.push.mock.calls[0]).toEqual([ '/register'])
}); 

test('User Listing update action', async () => {
  const mockEditListing = jest.fn();
  const mockItem = { id: 123, courseName: "Biology", content: "ILY", }
  render(
    <ShortListing item={mockItem} editListing={mockEditListing}/>
  );
  await fireEvent.click(screen.getByText('Update'));
  expect(mockEditListing.mock.calls[0]).toEqual([ mockItem.content, mockItem ])
}); 

test('User Listing delete action', async () => {
  const mockEditListing = jest.fn();
  const mockItem = { id: 123, courseName: "Biology", content: "ILY"}
  render(
    <ShortListing item={mockItem} editListing={mockEditListing}/>
  );
  await fireEvent.click(screen.getByText('Close'));
  expect(mockEditListing.mock.calls[0]).toEqual([ 'Close', mockItem ])
}); 

test('Course View has all course info', async () => {
  const mockRemoveCourse = jest.fn();
  const mockItem = { id: 123, name: "Biology", keywords: ["apple", "cookie", "tree"] }
  render(
    <CourseView item={mockItem} removeCourse={mockRemoveCourse}/>
  );
  const keywordsFormatted = await screen.findAllByText(/apple, cookie, tree/);
  const nameShown = await screen.findAllByText(/Biology/);
  expect(nameShown).toHaveLength(1);
  expect(keywordsFormatted).toHaveLength(1);
}); 

test('Course View remove course', async () => {
  const mockRemoveCourse = jest.fn();
  const mockItem = { id: 123, name: "Biology", keywords: ["apple", "cookie", "tree"] }
  render(
    <CourseView item={mockItem} removeCourse={mockRemoveCourse}/>
  );
  await fireEvent.click(screen.getByText('Remove Course'));
  expect(mockRemoveCourse.mock.calls[0]).toEqual([ mockItem ])
}); 

test('My Listings has all listings', async () => {
  const mockEditListing = jest.fn();
  const mockItems = [{ id: 121, courseName: "Biology", content: "ILY"}, { id: 123, courseName: "Biology", content: "ILY"}, { id: 124, courseName: "Biology", content: "ILY"}]
  render(
    <MyListings items={mockItems} editListing={mockEditListing}/>
  );
  const nameShown = await screen.findAllByText(/Biology/);
  expect(nameShown).toHaveLength(3);
}); 

test('Scheduler Page view has weekly 24 hr views', async () => {
  render(
    <SchedulerPage/>
  );
  const morningShown = await screen.findAllByText(/am/);
  const nightShown = await screen.findAllByText(/pm/);
  const daysShown = await screen.findAllByText(/day/);
  expect(morningShown).toHaveLength(168);
  expect(nightShown).toHaveLength(168);
  expect(daysShown).toHaveLength(7);
}); 

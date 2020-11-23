import React from 'react';
import { Route, Switch, MemoryRouter } from 'react-router-dom';
import { render, fireEvent, waitFor, screen, getByTestId, within, queryByTestId } from '@testing-library/react'
import '@testing-library/jest-dom/extend-expect'
import ClassItem from './components/CurrentClasses/ClassItem';
import GroupItem from './components/CurrentClasses/GroupItem';
import ClassList from './components/CurrentClasses/ClassList';
import Dropdown from './components/Dropdown/Dropdown';
import RequestRow from './components/Requests/RequestRow';
import UserItem from './components/UserList/UserItem';
import UserList from './components/Search/UserList';
import Inviter from './components/Inviter/Inviter';

test('Class Item Click', async () => {
  const mockClick = jest.fn();
  const mockData = { name: 'fakeName', clickable: true }
  render(
      <ClassItem data={mockData} titleClicked={mockClick}/>
  );
  fireEvent.click(screen.getByTestId('click-class'));
  expect(mockClick.mock.calls[0]).toEqual([ mockData ])
}); 

test('Group Item Click', async () => {
  const mockClick = jest.fn();
  const mockData = { name: 'fakeName' }
  render(
      <GroupItem group={mockData} titleClicked={mockClick}/>
  );
  fireEvent.click(screen.getByTestId('click-group'));
  expect(mockClick.mock.calls[0]).toEqual([ mockData ])
}); 

test('Class List renders all items', async () => {
  const mockClick = jest.fn();
  const mockData = [{ name: 'fakeName', clickable: true }, { name: 'fakeName', clickable: true }, { name: 'fakeName', clickable: true }]
  render(
      <ClassList classList={mockData} titleClicked={mockClick} clickable={true} />
  );
  const items = await screen.findAllByText(/fakeName/);
  expect(items).toHaveLength(3);
}); 

const mockEvent = { preventDefault: jest.fn() };
jest.mock('./components/Dropdown/DropdownRow', () => {
  return {
    __esModule: true,
    default: ({item, changeSelection}) => {
    return <button onClick={() => changeSelection(mockEvent, item)}>{item.name}</button>;
    },
  };
});

test('Dropdown row press shows dropdown', async () => {
  const mockData = [{ name: 'fakeName', clickable: true }, { name: 'fakeName2', clickable: true }]
  render(
      <Dropdown options={mockData} sendSelection={() => {}}/>
  );

  fireEvent.click(screen.getByRole('button'));
  const items = await screen.findAllByText(/fakeName/)
  expect(items).toHaveLength(3);
}); 

test('Link click works for view group in invitations', async () => {
  const mockData = { name: 'fakeName', type: 'invitation', groupId: 1};
  render(
    <MemoryRouter initialEntries={[`/`]}>
      <RequestRow item={mockData}/>
    </MemoryRouter>
  );
  await fireEvent.click(screen.getByText('fakeName'));
  expect(document.querySelector("a").getAttribute("href")).toBe(
    "/groups/group/1"
  );
}); 

test('Link click works for view profile in requests', async () => {
  const mockData = { name: 'fakeName', type: 'request', id: 1};
  render(
    <MemoryRouter initialEntries={[`/`]}>
      <RequestRow item={mockData}/>
    </MemoryRouter>
  );
  await fireEvent.click(screen.getByText('fakeName'));
  expect(document.querySelector("a").getAttribute("href")).toBe(
    "/profile/1"
  );
}); 

test('User does not have additional contact information, email shown', async () => {
  const mockUser = { name: "Edgar G", school: 'UCLA', email: "ed@gmail.com"};
  render(
    <UserItem user={mockUser} goToUserProfile={() => {}}/>
  );
  const items = await screen.findAllByText(/ed@gmail.com/)
  expect(items).toHaveLength(1);
}); 

test('User has additional contact information, email not shown', async () => {
  const mockUser = { name: "Edgar G", school: 'UCLA', email: "ed@gmail.com", discord:"ed#123"};
  render(
    <UserItem user={mockUser} goToUserProfile={() => {}}/>
  );
  const items = await screen.findAllByText(/ed#123/)
  const notFound = await screen.queryByText(/ed@gmail.com/)
  expect(items).toHaveLength(1);
  expect(notFound).not.toBeInTheDocument()
}); 

test('User row click leads to correct user profile', async () => {
  const mockClick = jest.fn();
  const mockUser = [{ id: 123, name: "Edgar G", school: 'UCLA', email: "ed@gmail.com", discord:"ed#123"}];
  render(
    <UserList userList={mockUser} goToUserProfile={mockClick}/>
  );

  fireEvent.click(screen.getByTestId('user'));
  expect(mockClick.mock.calls[0]).toEqual([ mockUser[0] ])
}); 

test('Click leads to invite user to first group', async () => {
  const mockClick = jest.fn();
  const mockUser = { id: 123, name: "Edgar G", school: 'UCLA', email: "ed@gmail.com", discord:"ed#123"};
  const mockGroup = [{id: 12, groupName: "Calc Crew"}, {id: 13, groupName: "Calc 123"}];
  render(
    <Inviter items={mockGroup} user={mockUser} handleGroupInvite={mockClick}/>
  );

  fireEvent.click(screen.getByText('Invite'));
  expect(mockClick.mock.calls[0]).toEqual([ mockUser, mockGroup[0] ]);
}); 
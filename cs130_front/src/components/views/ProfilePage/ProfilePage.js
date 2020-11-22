import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './styles.css';
import Requests from '../../Requests/Requests';
import Infos from '../../Infos/Infos';
import SimpleForm from '../../Infos/SimpleForm';
import ClassList from '../../CurrentClasses/ClassList';
import MyListings from './MyListings';
import SchedulerPage from '../SchedulerPage/SchedulerPage';
import CourseAdder from './CourseAdder';
import CourseView from './CourseView';
import * as Colors from '../../../constants/Colors';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import { USER_SERVER_AUTH , USER_SERVER } from '../../../Config';

function ProfilePage(props) {
  const [invitations, setInvitations] = useState([]);
  const [contactInfo, setContactInfo] = useState([]);
  const [additionalInfo, setAdditionalInfo] = useState([]);
  const [listings, setListings] = useState([]);
  const [classes, setClasses] = useState([]);
  const [myCourses, setMyCourses] = useState([]);
  const [availability, setAvailability] = useState([]);
  const [mainPanelState, setMainPanel] = useState('CourseAdder');
  const userId = props.match.params.id;
  const myId = localStorage.getItem('userId');

  const config = {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
    }
  }

  useEffect(() => {
		async function initProfile() {
			try {
        const response = await getProfile(userId);
        handleGetUserResponse(response);
        const classesInfoResponse = await getClassesInfo();
        handleClassesInfoResponse(classesInfoResponse.data.courses);
			} catch (err) {
				// Handle err here. Either ignore the error, or surface the error up to the user somehow.
			}
		}
		initProfile();
  }, []);
  
  async function getClassesInfo() {
    return axios.get(`${USER_SERVER}/classes-info`);
  }

  function getProfile() {
    return axios.get(`${USER_SERVER_AUTH}?u_id=${userId}`, config);
  }

  function handleGetUserResponse(response) {
    const contactInfoData = [
      { name: "Email", value: response.data.u_email }, 
      { name: "Facebook", value: response.data.facebook },
      { name: "Discord", value: response.data.discord }
    ];
    const additionalInfoData = [
      { name: "Name", value: `${response.data.first_name} ${response.data.last_name}` },
      { name: "School name", value: response.data.school_name }, 
      { name: "Timezone", value: response.data.timezone },
      { name: "Biography", value: response.data.biography },
    ];
    const fetchedListings = response.data.listings.map(listing => {
      return { id: listing.id, courseName: listing.course_name, content: listing.text_description }
    })
    const inviteData = response.data.invitations.map(invite => {
      return { name: invite.group_name, inviteId:invite.id, id: invite.receive_id, type: "invitation" };
    })
    const fetchedCourses = response.data.courses.map(course => {
      const groups = [];
      response.data.groups.forEach(group => {
        if (group.course_id === course.id) {
          groups.push(group);
        }
      })
      return { name: course.name, classId: course.id, keywords: course.keywords, groups: groups.length > 0 ? groups : null }
    })
    setContactInfo(contactInfoData);
    setAdditionalInfo(additionalInfoData);
    setInvitations(inviteData);
    setListings(fetchedListings);
    setMyCourses(fetchedCourses);
    setAvailability(response.data.availability);
  }

  function handleClassesInfoResponse(courses) {
    const classes = []
    Object.keys(courses).forEach(function(key) {
      const institution = key;
      const categories = [];
      Object.keys(courses[key]).forEach(function(key2) {
        const category = key2;
        const classes = courses[key][key2];
        categories.push({ category: category, classes: classes});
      });
      classes.push({ institution: institution, categories: categories })
    });
    setClasses(classes);
  }

  function saveInfoClicked(first, second, third) {
    var body;
    if(first.name == 'Email') {
      body = {
        u_id: parseInt(userId),
        u_email: first.value,
        discord: third.value,
        facebook: second.value, 
      }
    } else {
      const splitName = first.value.split(' ');
      body = {
        u_id: parseInt(userId),
        u_email: contactInfo[0].value,
        first_name: splitName[0],
        last_name: splitName[1],
        timezone: third.value,
        school_name: second.value,
      }
    }
    axios.put(`${USER_SERVER_AUTH}/update`, body, config).then(response => {
      return getProfile();
    }).then( getResponse => {
      handleGetUserResponse(getResponse);
    });
  }

  function handleInvitation(status, invitation) {
    console.log(status, invitation);
      const body = {
        u_id: parseInt(invitation.id),
        invitation_id: invitation.inviteId,
        status: status ? 'ACCEPT' : 'DECLINE'
      }
      axios.put(`${USER_SERVER_AUTH}/updateInvitation`, body, config).then(response => {
        return getProfile();
      }).then( getResponse => {
        handleGetUserResponse(getResponse);
      }); 
  }

  function editListing(content, listing) {
    console.log(content, listing);
    if(content == 'Close') {
      axios.delete(`${USER_SERVER_AUTH}/deleteListing?id=${listing.id}`, config).then(response => {
        return getProfile();
      }).then( getResponse => {
        handleGetUserResponse(getResponse);
      }); 
    } else {
      const body = {
        id: parseInt(listing.id),
        text_description: content
      }
      axios.put(`${USER_SERVER_AUTH}/updateListing`, body, config).then(response => {
        return getProfile();
      }).then( getResponse => {
        handleGetUserResponse(getResponse);
      }); 
    }
  }
  
  function addCourse(course) {

    const body = {
      u_id: parseInt(myId),
      course_id: parseInt(course.classId),
      course_name: course.name,
      keywords: course.keywords instanceof String ? [course.keywords] : course.keywords,
      categories: course.categories

    }
    axios.put(`${USER_SERVER_AUTH}/addCourse`, body, config).then(response => {
      return getProfile();
    }).then( getResponse => {
      handleGetUserResponse(getResponse);
    }); 
  }

  function removeCourse(course) {
    const body = {
      u_id: parseInt(myId),
      course_id: parseInt(course.classId),
    }
    axios.put(`${USER_SERVER_AUTH}/removeCourse`, body, config).then(response => {
      setMainPanel('CourseAdder');
      return getProfile();
    }).then( getResponse => {
      handleGetUserResponse(getResponse);
    }); 
  }

  function renderMainPanel() {
    switch (mainPanelState) {
      case 'Contact Information':
        return <SimpleForm key="contact" options={contactInfo} saveInfoClicked={saveInfoClicked}/>
      case 'Additional Information': 
        return <SimpleForm key="additional" options={additionalInfo} saveInfoClicked={saveInfoClicked}/>
      case 'CourseAdder': 
        return <CourseAdder courses={classes} addCourse={addCourse}/>
      default: //Course view for selected course
        const course = myCourses.find(element => element.name == mainPanelState.name)
        return <CourseView item={course} removeCourse={removeCourse}/>
    }
  }

  function coursesOnly() {
    var newCourses = [...myCourses];
    newCourses.forEach( element => {
      delete element["groups"]
    })
    return newCourses;
  }

  function titleClicked(title) {
    //set main content to be for title
    setMainPanel(title);
  }

  function groupClicked(item) {
    //navigate to group page
    console.log(item);
  }

  function viewOnlyProfile() {
    return (
      <div className="view-panel">
        <div className="column"> 
          <div className="text-container">
            <Text color="black" size="32px" weight="800"> 
            {additionalInfo[0].value} Profile
            </Text>
            <ClassList classList={userId == myId ? coursesOnly() : myCourses} titleClicked={groupClicked} clickable={false}/>
            <Infos title="Contact Information" options={contactInfo} titleClicked={titleClicked} clickable={false}/>
            <Infos title="Additional Information" options={additionalInfo} titleClicked={titleClicked} clickable={false}/>
          </div>
        </div>
        <div className="column">
          <SchedulerPage passedSelections={availability}/>
        </div>
      </div>
    );
  }
  
  function myProfile() {
    return (
      <div className="panel">
  
        <div className="column"> 
          <div className="text-container">
            <Text color="black" size="24px" weight="800"> 
              Hi Student, edit your class and other information to start studying with others!
            </Text>
          </div>
          <MyListings items={listings} editListing={editListing}/> 
          { invitations.length > 0 ? 
            <Requests title="Invitations" items={invitations} handleResponse={handleInvitation}/> : null
          }
        </div>
        <div className="column"> 
          {renderMainPanel()}
        </div>
        <div className="column">
          <div className="group-with-margin-bottom">
            <ClassList classList={userId == myId ? coursesOnly() : myCourses} titleClicked={titleClicked} clickable={true}/>
            <Infos title="Contact Information" options={contactInfo} titleClicked={titleClicked} clickable={true}/>
            <Infos title="Additional Information" options={additionalInfo} titleClicked={titleClicked} clickable={true}/>
          </div>
          <div className="group-with-margin-bottom">
            <Button 
            textColor={Colors.White}
            textSize="28px"
            width="280px"
            height="70px"
            textWeight="800" 
            color={Colors.Blue}
            onClick={() => setMainPanel('CourseAdder')}
            >
              Add Courses
            </Button>
          </div>
          <Button 
          textColor={Colors.White}
          textSize="28px"
          width="280px"
          height="70px"
          textWeight="800" 
          color={Colors.Blue}
          onClick={() => props.history.push({
            pathname: `/profile/${userId}/scheduler`,
            state: { availability: availability, email: contactInfo[0].value }
            })}
          >
            Set Availability
          </Button>
        </div>
      </div>
    );
  }

  // will have conditional logic based on self profile vs different.. maybe a diff endpoint?
  if (additionalInfo.length > 0 && classes.length > 0) {
    return userId == myId ? myProfile() : viewOnlyProfile();
  }
  return null;
}

export default ProfilePage;

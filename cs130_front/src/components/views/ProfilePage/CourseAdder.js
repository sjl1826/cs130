import React, {useState, useEffect } from 'react';
import Dropdown from '../../Dropdown/Dropdown';
import Button from '../../Button/Button';
import Text from '../../Text/Text';
import * as Colors from '../../../constants/Colors';
import CourseCreater from './CourseCreater';
import './styles.css';

export default function CourseAdder(props) {
  const mappedInstitutions = props.courses.map(institution => {
    return {
      name: institution.institution
    }
  });
  mappedInstitutions.unshift({name: "Choose an Institution first"});
  
  const [selectedInstitution, setInstitution] = useState('');
  const [selectedCategory, setCategory] = useState(null);
  const [selectedCourse, setCourse] = useState(null);
  const [selectedCategories, setCategories] = useState(null);
  const [classes, setClasses] = useState(null);

  function selectInstitution(item) {
    if(item.name == "Choose an Institution first" || selectedInstitution == item.name) {
      return;
    }
    setInstitution(item.name);
    setCategory(null);
    setCourse(null);
    setClasses(null);
    setCategories(null);
  } 

  function selectCategory(item) {
    if(item.name == "Choose a category" || selectedCategory == item.name) {
      return;
    }
    setCategory(item.name);
    setCourse(null);
    setClasses(null);
  } 

  function selectCourse(item) {
    setCourse(item);
  }

  function renderCourseOrCreate(){
    if (selectedCourse != null && selectedCourse.name == "Create a course") {
      return(<CourseCreater addCustomCourse={addCustomCourse}/>);
    }

    return (selectedCourse != null ? <Text>Keywords: {selectedCourse.keywords.join(', ')}</Text> : null);
  }

  function addCustomCourse(name, keywords) {
    const categories = [selectedInstitution, selectedCategory, name];
    const course = {name: name, keywords: [keywords], classId: 0, categories: categories}
    setCourse(course);
    props.addCourse(course)
  }

  useEffect(() => {
    props.courses.forEach(each => {
      if (each.institution == selectedInstitution) {
        const mappedCategories = each.categories.map(category => {
          return {
            name: category.category
          }
        });
        mappedCategories.unshift({name: "Choose a category"});
        setCategories(mappedCategories);
      }
    });
  }, [selectedInstitution]);

  useEffect(() => {
    props.courses.forEach(each => {
      if (each.institution == selectedInstitution) {
        each.categories.forEach(category => {
          if (category.category == selectedCategory) {
            const mappedCourses = category.classes.map(course => {
              return {
                name: course.name,
                keywords: course.keywords,
                classId: course.id,
                categories: course.categories
              }
            });
            mappedCourses.push({name: "Create a course"});
            setClasses(mappedCourses);
          }
        })
      }
    });
  }, [selectedCategory]);


  return (
    <div className="formTwo">
      <div className="row">
        <Dropdown width="15vw"options={mappedInstitutions} sendSelection={selectInstitution}/>
        {selectedInstitution != null && selectedCategories != null ? <Dropdown width="15vw" options={selectedCategories} sendSelection={selectCategory}/> : null}
        {selectedCategories != null && classes != null ? <Dropdown width="15vw" options={classes} sendSelection={selectCourse}/> : null}
      </div>
      {renderCourseOrCreate()}
      {selectedCourse != null && selectedCourse.name == 'Create a course' ? null :       
        <Button 
          textColor={Colors.White}
          textSize="28px"
          width="280px"
          height="70px"
          textWeight="800" 
          color={Colors.Blue}
          onClick={() => props.addCourse(selectedCourse)}
        >
          Add Course
        </Button>
      }
    </div>
  );
}

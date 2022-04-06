![angularjs_logo](https://user-images.githubusercontent.com/4659608/37036392-9bf53686-2160-11e8-95fc-bbab638d7d60.png)

# Angular-csv-ext | Export to CSV  in Angular


[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e2133aa828054d7c865563b50100eb8b)](https://www.codacy.com/app/me_101/angular-csv-ext?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=alhazmy13/angular-csv-ext&amp;utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/alhazmy13/angular-csv-ext.svg?branch=master)](https://travis-ci.org/alhazmy13/angular-csv-ext)
[![npm version](https://badge.fury.io/js/angular-csv-ext.svg)](https://badge.fury.io/js/angular-csv-ext)
[![GitHub license](https://img.shields.io/github/license/alhazmy13/angular-csv-ext.svg)](https://github.com/alhazmy13/angular-csv-ext)
![Angular](https://img.shields.io/badge/Angular-%3E%3D5.0-red.svg)
![npm](https://img.shields.io/npm/dm/angular-csv-ext.svg)

> A helper library for creating CSV files in Angular.
> 

## Installation 

```javascript
npm install --save angular-csv-ext
```

## Example 
```javascript

import { AngularCsv } from 'angular-csv-ext/dist/Angular-csv';

var data = [
  {
    name: "Test 1",
    age: 13,
    average: 8.2,
    approved: true,
    description: "using 'Content here, content here' "
  },
  {
    name: 'Test 2',
    age: 11,
    average: 8.2,
    approved: true,
    description: "using 'Content here, content here' "
  },
  {
    name: 'Test 4',
    age: 10,
    average: 8.2,
    approved: true,
    description: "using 'Content here, content here' "
  },
];

new AngularCsv(data, 'My Report');

```

## API | **AngularCsv(data, filename, options)**


| Option        | Default           | Description  |
| :------------- |:-------------:| -----|
| **fieldSeparator**      | , | Defines the field separator character |
| **quoteStrings**      | "      | If provided, will use this characters to "escape" fields, otherwise will use double quotes as deafult |
| **decimalseparator** | .      | Defines the decimal separator character (default is .). If set to "locale", it uses the [language sensitive representation of the number](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/toLocaleString).|
| **showLabels** | false      | If provided, would use this attribute to create a header row |
| **showTitle** | false      |   |
| **useBom** | true      | If true, adds a BOM character at the start of the CSV |
| **useHeader** | false      | If true, only fields listed in header will be exported in CSV |
| **noDownload** | false      | If true, disables automatic download and returns only formatted CSV |
| **nullToEmptyString** | false      | If true, all null values will be changed to empty strings |


## Options Example

```javascript
  var options = { 
    fieldSeparator: ',',
    quoteStrings: '"',
    decimalseparator: '.',
    showLabels: true, 
    showTitle: true,
    title: 'Your title',
    useBom: true,
    noDownload: true,
    headers: ["First Name", "Last Name", "ID"],
    useHeader: false,
    nullToEmptyString: true,
  };

  AngularCsv(data, filename, options);

```

## Credits


 * [sn123](https://github.com/sn123)
 * [arf1980](https://github.com/arf1980)

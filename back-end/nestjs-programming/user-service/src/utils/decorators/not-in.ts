import {
  registerDecorator,
  ValidationArguments,
  ValidationOptions,
} from 'class-validator';

// class-validator 에서 제공하는 데코레이터와 같은 형태로 인자를 받음
// 체크할 프로퍼티 이름 및 커스텀할 수 있는 ValidationOptions 인자
export function NotIn(property: string, validationOptions?: ValidationOptions) {
  // 검증할 객체(object) 및 프로퍼티(propertyName) 정보를 인자로 받는 함수
  return (object: object, propertyName: string) => {
    // class-validator 에 커스텀 validator decorator 등록
    registerDecorator({
      name: 'NotIn',
      target: object.constructor,
      propertyName,
      options: validationOptions,
      constraints: [property],
      validator: {
        // ValidatorConstraintInterface 구현
        validate(
          value: any,
          validationArguments?: ValidationArguments,
        ): Promise<boolean> | boolean {
          // NotIn 매개변수의 property 키에 해당하는 값(relatedValue)
          const [relatedPropertyName] = validationArguments.constraints;
          const relatedValue = (validationArguments.object as any)[
            relatedPropertyName
          ];

          // 데코레이터 인자로 넘어 온 프로퍼티 값은 데코레이터가 적용된 프로퍼티 값을 포함하면 안됨
          return (
            typeof value === 'string' &&
            typeof relatedValue === 'string' &&
            !relatedValue.includes(value)
          );
        },
      },
    });
  };
}
